package jsonrpc2

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ToArray(params json.RawMessage, expected int) ([]json.RawMessage, error) {
	var rawArgs []json.RawMessage
	err := json.Unmarshal(params, &rawArgs)
	if err != nil {
		return nil, fmt.Errorf("parse arguments: %v", err)
	}
	if len(rawArgs) != expected {
		return nil, fmt.Errorf("missmatch number of arguments: expected %d but got %d", expected, len(rawArgs))
	}
	return rawArgs, nil
}

func UnmarshalArray(params json.RawMessage, args ...interface{}) error {
	array, err := ToArray(params, len(args))
	if err != nil {
		return err
	}
	for i, v := range args {
		err := json.Unmarshal(array[i], v)
		if err != nil {
			return fmt.Errorf("parse #%d: %v", i, err)
		}
	}
	return nil
}

// Expose JSON-RPC router as HTTP handler. Supported methods: POST, PUT, PATCH
func Handler(router *Router) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()
		switch request.Method {
		case http.MethodPost, http.MethodPut, http.MethodPatch:
		default:
			http.Error(writer, "Method not supported", http.StatusMethodNotAllowed)
			return
		}
		resp, isBatch := router.Invoke(request.Body)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(writer)
		enc.SetIndent("", "  ")
		var err error
		if isBatch {
			err = enc.Encode(resp)
		} else if len(resp) > 0 {
			err = enc.Encode(resp[0])
		}
		if err != nil {
			log.Println("server error:", err)
		}
	}
}

// Interceptor that limiting maximum number of requests in a batch. If batch size is bigger
//- InternalError will be returned with description. Only one error response without ID will be generated regardless of batch size
func MaxBatch(num int) GlobalInterceptorFunc {
	return func(gic *GlobalInterceptorContext) (responses []*Response, isBatch bool) {
		if len(gic.Requests) > num {
			return []*Response{
				{
					Version: Version,
					Error: &Error{
						Code:    InternalError,
						Message: "batch is too big",
					},
				},
			}, gic.IsBatch
		}
		return gic.Next()
	}
}
