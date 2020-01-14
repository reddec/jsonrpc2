package jsonrpc2

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"sync"
)

// Router for JSON-RPC requests.
//
// Supports batching.
type Router struct {
	methods map[string]Method
	lock    sync.RWMutex
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

// Invoke exposed method using request from stream (as a batch or single)
func (caller *Router) Invoke(stream io.Reader) (responses []*Response, isBatch bool) {
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
	//TODO: global hooks
	//TODO: per method hook

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
				responses[i] = request.failed(InvalidRequest, "Invalid Request")
				return
			}
			caller.lock.RLock()
			invoker, ok := caller.methods[request.Method]
			caller.lock.RUnlock()
			if !ok {
				responses[i] = request.failed(MethodNotFound, "Method not found")
				return
			}
			isPositional := len(request.Params) > 0 && request.Params[0] == '['
			reply, err := invoker.JsonCall(request.Params, isPositional)
			if err != nil {
				responses[i] = request.failed(AppError, err.Error())
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
