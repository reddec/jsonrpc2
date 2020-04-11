package jsonrpc2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
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

// Expose JSON-RPC route over HTTP Rest (POST) and web sockets (GET)
func Handler(router *Router) http.HandlerFunc {
	rest := HandlerRest(router)
	ws := HandlerWS(router)
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			ws.ServeHTTP(writer, request)
		} else {
			rest.ServeHTTP(writer, request)
		}
	}
}

// Expose JSON-RPC router as HTTP handler where one requests is one execution.
// Supported methods: POST, PUT, PATCH
func HandlerRest(router *Router) http.HandlerFunc {
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

// Process requests over web socket (all requests are processing in parallel in a separate go-routine)
func HandlerWS(router *Router) http.HandlerFunc {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  8192,
		WriteBufferSize: 8192,
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()

		conn, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		var writeSync sync.Mutex
		var wg sync.WaitGroup
		for {
			msgType, p, err := conn.ReadMessage()
			if err != nil {
				break
			}
			if msgType != websocket.TextMessage {
				continue
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				resp, isBatch := router.Invoke(bytes.NewReader(p))
				var replyData []byte
				if isBatch {
					replyData, err = json.MarshalIndent(resp, "", "  ")
				} else if len(resp) > 0 {
					replyData, err = json.MarshalIndent(resp[0], "", "  ")
				}
				if err != nil {
					log.Println("server error:", err)
					return
				}
				writeSync.Lock()
				_ = conn.WriteMessage(websocket.TextMessage, replyData)
				writeSync.Unlock()
			}()
		}
		wg.Wait()
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
