package jsonrpc2

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"sync"
)

// Router for JSON-RPC requests.
//
// Supports batching.
type Router struct {
	methods     map[string]Method
	lock        sync.RWMutex
	methodHooks struct {
		listeners []MethodInterceptorFunc
		lock      sync.RWMutex
	}
	globalHooks struct {
		listeners []GlobalInterceptorFunc
		lock      sync.RWMutex
	}
}

// Register method to router to expose over JSON-RPC interface
func (caller *Router) Register(method string, handler Method) *Router {
	caller.lock.Lock()
	defer caller.lock.Unlock()
	if caller.methods == nil {
		caller.methods = make(map[string]Method)
	}
	caller.methods[method] = handler
	return caller
}

// Register function as method to expose over JSON-RPC
func (caller *Router) RegisterFunc(method string, handlerFunc MethodFunc) *Router {
	return caller.Register(method, handlerFunc)
}

// Register function as exposed method. Handler must return two values, last of them - error.
//
// For such methods only positional arguments supported.
func (caller *Router) RegisterPositionalOnly(method string, handler interface{}) error {
	wrapper, err := Function(handler)
	if err != nil {
		return err
	}
	caller.Register(method, wrapper)
	return nil
}

// Register function as exposed method. Function handler must have first argument is pointer to structure
// and must return payload and error.
//
// This kind of wrapper supports only named arguments.
func (caller *Router) RegisterNamedOnly(method string, handler interface{}) error {
	wrapper, err := RPCLike(handler)
	if err != nil {
		return err
	}
	caller.Register(method, wrapper)
	return nil
}

// Add interceptor for handling all methods invoke. Called in a same thread as method
func (caller *Router) InterceptMethods(handler MethodInterceptorFunc) *Router {
	caller.methodHooks.lock.Lock()
	caller.methodHooks.listeners = append(caller.methodHooks.listeners, handler)
	caller.methodHooks.lock.Unlock()
	return caller
}

// Add interceptor for handling batch before lookup for methods and execution
func (caller *Router) Intercept(handler GlobalInterceptorFunc) *Router {
	caller.globalHooks.lock.Lock()
	caller.globalHooks.listeners = append(caller.globalHooks.listeners, handler)
	caller.globalHooks.lock.Unlock()
	return caller
}

// Invoke exposed method using request from stream (as a batch or single) with background context
func (caller *Router) Invoke(stream io.Reader) (responses []*Response, isBatch bool) {
	return caller.InvokeContext(context.Background(), stream)
}

// Invoke exposed method using request from stream (as a batch or single) with custom context
func (caller *Router) InvokeContext(ctx context.Context, stream io.Reader) (responses []*Response, isBatch bool) {
	var batch []*Request
	var toUnparse interface{} = &batch
	var bufferedStream = bufio.NewReader(stream)
	head, _ := bufferedStream.Peek(1)
	if bytes.Compare([]byte{'['}, head) != 0 {
		first := &Request{}
		batch = append(batch, first)
		toUnparse = first
	} else {
		isBatch = true
	}

	err := json.NewDecoder(bufferedStream).Decode(toUnparse)
	if err != nil {
		responses = append(responses, &Response{
			Version: Version,
			ID:      nil,
			Error: &Error{
				Code:    ParseError,
				Message: err.Error(),
			},
		})
		return
	}
	// global hooks
	return caller.callWithGlobalInterceptors(ctx, batch, isBatch)
}

func (caller *Router) findMethod(name string) (Method, bool) {
	caller.lock.RLock()
	invoker, ok := caller.methods[name]
	caller.lock.RUnlock()
	return invoker, ok
}

func (caller *Router) callWithGlobalInterceptors(ctx context.Context, requests []*Request, isBatch bool) ([]*Response, bool) {
	caller.globalHooks.lock.RLock()
	defer caller.globalHooks.lock.RUnlock()
	gic := &GlobalInterceptorContext{
		Requests: requests,
		IsBatch:  isBatch,
		Context:  ctx,
		idx:      0,
		list:     caller.globalHooks.listeners,
		caller:   caller,
	}
	return gic.Next()
}

func (caller *Router) callWithMethodsInterceptors(ctx context.Context, method Method, request *Request, isPositional bool) (interface{}, error) {
	if len(caller.methodHooks.listeners) == 0 {
		return method.JsonCall(ctx, request.Params, isPositional)
	}
	caller.methodHooks.lock.RLock()
	defer caller.methodHooks.lock.RUnlock()
	ic := &MethodInterceptorContext{
		Request:      request,
		IsPositional: isPositional,
		Context:      ctx,
		list:         caller.methodHooks.listeners,
		method:       method,
	}
	return ic.Next()
}

// Context handling request per method
type MethodInterceptorContext struct {
	Request      *Request
	IsPositional bool
	Context      context.Context
	idx          int
	list         []MethodInterceptorFunc
	method       Method
}

// Call next interceptor or final method
func (ic *MethodInterceptorContext) Next() (interface{}, error) {
	if ic.idx >= len(ic.list) {
		return ic.method.JsonCall(ic.Context, ic.Request.Params, ic.IsPositional)
	}
	idx := ic.idx
	ic.idx++
	return ic.list[idx](ic)
}

// Interceptor for each method that will be called
type MethodInterceptorFunc func(ic *MethodInterceptorContext) (interface{}, error)

type GlobalInterceptorFunc func(gic *GlobalInterceptorContext) (responses []*Response, isBatch bool)

type GlobalInterceptorContext struct {
	Requests []*Request
	Context  context.Context
	IsBatch  bool
	idx      int
	list     []GlobalInterceptorFunc
	caller   *Router
}

func (gic *GlobalInterceptorContext) Next() (responses []*Response, isBatch bool) {
	if gic.idx >= len(gic.list) {
		return gic.invoke()
	}
	idx := gic.idx
	gic.idx++
	return gic.list[idx](gic)
}

func (gic *GlobalInterceptorContext) invoke() (responses []*Response, isBatch bool) {
	var batch = gic.Requests
	isBatch = gic.IsBatch
	var numNotifications = 0
	for _, invoke := range batch {
		if invoke.IsNotification() {
			numNotifications++
		}
	}
	responses = make([]*Response, len(batch))
	// invoke all
	wg := sync.WaitGroup{}
	for i, request := range batch {
		wg.Add(1)
		go func(i int, request *Request) {
			defer wg.Done()
			if !request.IsValid() {
				responses[i] = request.failed(InvalidRequest, "Invalid Request", nil)
				return
			}
			invoker, ok := gic.caller.findMethod(request.Method)
			if !ok {
				responses[i] = request.failed(MethodNotFound, "Method not found", nil)
				return
			}
			isPositional := len(request.Params) > 0 && request.Params[0] == '['

			// per method hook
			reply, err := gic.caller.callWithMethodsInterceptors(gic.Context, invoker, request, isPositional)
			var jsonRpcErr *Error
			if errors.As(err, &jsonRpcErr) {
				responses[i] = request.failed(jsonRpcErr.Code, jsonRpcErr.Message, jsonRpcErr.Data)
				return
			} else if err != nil {
				responses[i] = request.failed(AppError, err.Error(), nil)
				return
			}
			responses[i] = request.success(reply)
		}(i, request)
	}
	wg.Wait()
	if numNotifications > 0 {
		// remove notifications responses
		filtered := make([]*Response, 0, len(batch)-numNotifications)
		for i, res := range responses {
			if !batch[i].IsNotification() {
				filtered = append(filtered, res)
			}
		}
		responses = filtered
	}
	return
}
