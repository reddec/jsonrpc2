package jsonrpc2

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

const (
	Version        = "2.0"
	AppError       = -30000
	ParseError     = -32700
	InvalidRequest = -32600
	MethodNotFound = -32601
	InvalidParams  = -32602
	InternalError  = -32603
)

// Standard JSON-RPC 2.0 request messages
type Request struct {
	// always 2.0 (will refuse if not)
	Version string `json:"jsonrpc"`
	// case-sensitive method name
	Method string `json:"method"`
	// any kind of valid JSON as ID (more relaxed comparing to for spec)
	ID json.RawMessage `json:"id"`
	// array (for positional) or object (for named) of arguments
	Params json.RawMessage `json:"params"`
}

// Base checks against specification
func (rq *Request) IsValid() bool {
	if rq.Version != Version {
		return false
	}
	return true
}

// Check request is notification (null ID)
func (rq *Request) IsNotification() bool { return rq.ID == nil }

func (rq *Request) failed(code int, message string) *Response {
	return &Response{
		Version: Version,
		ID:      rq.ID,
		Error: &Error{
			Code:    code,
			Message: message,
		},
	}
}

func (rq *Request) success(payload interface{}) *Response {
	return &Response{
		Version: Version,
		ID:      rq.ID,
		Result:  payload,
	}
}

// JSON-RPC 2.0 standard error object
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// JSON-RPC 2.0 standard response object
type Response struct {
	// always 2.0
	Version string `json:"jsonrpc"`
	// any kind of valid JSON as ID (more relaxed comparing to for spec) copied from request
	ID json.RawMessage `json:"id"`
	// result if exists
	Result interface{} `json:"result,omitempty"`
	// error if exists
	Error *Error `json:"error,omitempty"`
}

// Method handler (for low-level implementation). Should support params as object or as array (positional=true).
//
// Returned data should be JSON serializable and not nil for success
type Method interface {
	JsonCall(params json.RawMessage, positional bool) (interface{}, error)
}

type MethodFunc func(params json.RawMessage, positional bool) (interface{}, error)

func (m MethodFunc) JsonCall(params json.RawMessage, positional bool) (interface{}, error) {
	return m(params, positional)
}

// Wrap function as JSON-RPC method for usage in router
//
// This kind of wrapper support only positional arguments
func Function(handler interface{}) (*callableWrapper, error) {
	val := reflect.ValueOf(handler)
	if val.IsNil() {
		return nil, errors.New("handler is nil")
	}
	typ := val.Type()
	if typ.Kind() != reflect.Func {
		return nil, errors.New("handler is not a function")
	}
	if typ.NumOut() != 2 {
		return nil, errors.New("function should return exactly two values: payload and error")
	}
	errorInterface := reflect.TypeOf((*error)(nil)).Elem()
	if !typ.Out(1).Implements(errorInterface) {
		return nil, errors.New("last return value of function should be an error")
	}
	return &callableWrapper{tp: typ, fn: val}, nil
}

type callableWrapper struct {
	fn reflect.Value
	tp reflect.Type
}

func (c *callableWrapper) JsonCall(params json.RawMessage, positional bool) (interface{}, error) {
	if !positional {
		return nil, errors.New("exported function supports only positional arguments")
	}
	var rawArgs []json.RawMessage
	err := json.Unmarshal(params, &rawArgs)
	if err != nil {
		return nil, fmt.Errorf("parse positional arguments: %v", err)
	}

	N := c.tp.NumIn()
	if len(rawArgs) != N {
		return nil, fmt.Errorf("mismatch number of arguments: expected %v but got %v", N, len(rawArgs))
	}

	var args = make([]reflect.Value, N)

	for i := 0; i < N; i++ {
		tp := c.tp.In(i)
		v := reflect.New(tp)

		rawArg := rawArgs[i]
		err = json.Unmarshal(rawArg, v.Interface())
		if err != nil {
			return nil, fmt.Errorf("parse arg %d: %v", i, err)
		}
		args[i] = v.Elem()
	}

	res := c.fn.Call(args)
	var outErr error
	if !res[1].IsNil() {
		outErr = res[1].Interface().(error)
	}
	return res[0].Interface(), outErr
}

// Expose function handler where first argument is pointer to structure and returns are payload with error.
//
// This kind of wrapper support only named arguments
func RPCLike(handler interface{}) (*rpcLikeCallable, error) {
	val := reflect.ValueOf(handler)
	if val.IsNil() {
		return nil, errors.New("handler is nil")
	}
	typ := val.Type()
	if typ.Kind() != reflect.Func {
		return nil, errors.New("handler is not a function")
	}
	if typ.NumOut() != 2 {
		return nil, errors.New("function should return exactly two values: payload and error")
	}
	errorInterface := reflect.TypeOf((*error)(nil)).Elem()
	if !typ.Out(1).Implements(errorInterface) {
		return nil, errors.New("last return value of function should be an error")
	}
	if typ.NumIn() != 1 {
		return nil, errors.New("number of input arguments should be exactly one")
	}
	tp := typ.In(0)
	if tp.Kind() != reflect.Ptr || tp.Elem().Kind() != reflect.Struct {
		return nil, errors.New("first argument should be pointer to struct")
	}

	return &rpcLikeCallable{argType: tp, fn: val, fnType: typ}, nil
}

type rpcLikeCallable struct {
	argType reflect.Type
	fn      reflect.Value
	fnType  reflect.Type
}

func (r *rpcLikeCallable) JsonCall(params json.RawMessage, positional bool) (interface{}, error) {
	if positional {
		return nil, errors.New("exported function supports only named arguments")
	}
	return r.callByNamed(params)
}

func (r *rpcLikeCallable) callByNamed(params json.RawMessage) (interface{}, error) {
	arg := reflect.New(r.argType)
	err := json.Unmarshal(params, arg.Interface())
	if err != nil {
		return nil, fmt.Errorf("failed to parse arguments: %v", err)
	}
	res := r.fn.Call([]reflect.Value{arg.Elem()})
	var outErr error
	if !res[1].IsNil() {
		outErr = res[1].Interface().(error)
	}
	return res[0].Interface(), outErr
}
