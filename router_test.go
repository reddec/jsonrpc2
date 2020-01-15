package jsonrpc2

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync/atomic"
	"testing"
)

func TestPositional(t *testing.T) {
	router := &Router{}
	err := router.RegisterPositionalOnly("sum", func(a, b int, c *int) (int, error) {
		return a + b + *c, nil
	})
	if err != nil {
		t.Error("reg function", err)
		return
	}

	responses, isBatch := router.Invoke(bytes.NewBufferString(`{
		"jsonrpc":"2.0",
		"id": 1,
		"method":"sum",
		"params": [1, 2, 3]
	}`))
	if isBatch {
		t.Error("should be not a batch")
		return
	}
	if len(responses) != 1 {
		t.Errorf("only one response expected but got %v", len(responses))
		return
	}
	if bytes.Compare(responses[0].ID, []byte("1")) != 0 {
		t.Errorf("unmatched ID: got %v", string(responses[0].ID))
	}
	if responses[0].Version != "2.0" {
		t.Errorf("version is not 2.0: %v", responses[0].Version)
	}
	if responses[0].Error != nil {
		t.Errorf("failed: %+v", responses[0].Error)
		return
	}
	if v, ok := responses[0].Result.(int); !ok || v != 6 {
		t.Errorf("not matched result: %v", responses[0].Result)
	}
}

func TestNamed(t *testing.T) {
	router := &Router{}
	err := router.RegisterNamedOnly("sum", func(params *struct {
		A int  `json:"a"`
		B int  `json:"b"`
		C *int `json:"c"`
	}) (int, error) {
		return params.A + params.B + *params.C, nil
	})
	if err != nil {
		t.Error("reg function", err)
		return
	}

	responses, isBatch := router.Invoke(bytes.NewBufferString(`{
		"jsonrpc":"2.0",
		"id": 1,
		"method":"sum",
		"params": {"a":1, "b":2, "c":3}
	}`))
	if isBatch {
		t.Error("should be not a batch")
		return
	}
	if len(responses) != 1 {
		t.Errorf("only one response expected but got %v", len(responses))
		return
	}
	if bytes.Compare(responses[0].ID, []byte("1")) != 0 {
		t.Errorf("unmatched ID: got %v", string(responses[0].ID))
	}
	if responses[0].Version != "2.0" {
		t.Errorf("version is not 2.0: %v", responses[0].Version)
	}
	if responses[0].Error != nil {
		t.Errorf("failed: %+v", responses[0].Error)
		return
	}
	if v, ok := responses[0].Result.(int); !ok || v != 6 {
		t.Errorf("not matched result: %v", responses[0].Result)
	}
}

func TestBatched(t *testing.T) {
	router := &Router{}
	err := router.RegisterPositionalOnly("sum", func(a, b int, c *int) (int, error) {
		return a + b + *c, nil
	})
	if err != nil {
		t.Error("reg function", err)
		return
	}

	responses, isBatch := router.Invoke(bytes.NewBufferString(`[{
		"jsonrpc":"2.0",
		"id": 1,
		"method":"sum",
		"params": [1, 2, 3]
	},
	{
		"jsonrpc":"2.0",
		"id": 2,
		"method":"sum",
		"params": [4, 5, 6]
	}]`))
	if !isBatch {
		t.Error("should be  batch")
		return
	}
	if len(responses) != 2 {
		t.Errorf("only two response expected but got %v", len(responses))
		return
	}
	if bytes.Compare(responses[0].ID, []byte("1")) != 0 {
		t.Errorf("(0) unmatched ID: got %v", string(responses[0].ID))
	}
	if bytes.Compare(responses[1].ID, []byte("2")) != 0 {
		t.Errorf("(1) unmatched ID: got %v", string(responses[1].ID))
	}
	if responses[0].Version != "2.0" {
		t.Errorf("(0) version is not 2.0: %v", responses[0].Version)
	}
	if responses[1].Version != "2.0" {
		t.Errorf("(1) version is not 2.0: %v", responses[1].Version)
	}
	if responses[0].Error != nil {
		t.Errorf("(0) failed: %+v", responses[0].Error)
		return
	}
	if responses[1].Error != nil {
		t.Errorf("(1) failed: %+v", responses[1].Error)
		return
	}
	if v, ok := responses[0].Result.(int); !ok || v != 6 {
		t.Errorf("(0) not matched result: %v", responses[0].Result)
	}
	if v, ok := responses[1].Result.(int); !ok || v != 15 {
		t.Errorf("(1) not matched result: %v", responses[1].Result)
	}
}

func TestNotification_single(t *testing.T) {
	router := &Router{}
	var invoked bool
	err := router.RegisterPositionalOnly("sum", func(a, b int, c *int) (int, error) {
		invoked = true
		return a + b + *c, nil
	})
	if err != nil {
		t.Error("reg function", err)
		return
	}

	responses, isBatch := router.Invoke(bytes.NewBufferString(`{
		"jsonrpc":"2.0",
		"method":"sum",
		"params": [1, 2, 3]
	}`))
	if isBatch {
		t.Error("should be not a batch")
		return
	}
	if len(responses) != 0 {
		t.Errorf("should be not responses but got %v", len(responses))
		return
	}
	if !invoked {
		t.Errorf("method not invoked")
	}
}

func TestNotification_batch(t *testing.T) {
	router := &Router{}
	var invoked int32
	err := router.RegisterPositionalOnly("sum", func(a, b int, c *int) (int, error) {
		atomic.AddInt32(&invoked, 1)
		return a + b + *c, nil
	})
	if err != nil {
		t.Error("reg function", err)
		return
	}

	responses, isBatch := router.Invoke(bytes.NewBufferString(`[{
		"jsonrpc":"2.0",
		"method":"sum",
		"params": [1, 2, 3]
	},
	{
		"jsonrpc":"2.0",
		"method":"sum",
		"params": [4, 5, 6]
	}]`))
	if !isBatch {
		t.Error("should be a batch")
		return
	}
	if len(responses) != 0 {
		t.Errorf("should be not responses but got %v", len(responses))
		return
	}
	if invoked != 2 {
		t.Errorf("method not invoked enouhg times: %v but expected 2 times", invoked)
	}
}

func TestNotificationWithRegular_batch(t *testing.T) {
	router := &Router{}
	var invoked int32
	err := router.RegisterPositionalOnly("sum", func(a, b int, c *int) (int, error) {
		atomic.AddInt32(&invoked, 1)
		return a + b + *c, nil
	})
	if err != nil {
		t.Error("reg function", err)
		return
	}

	responses, isBatch := router.Invoke(bytes.NewBufferString(`[{
		"jsonrpc":"2.0",
		"method":"sum",
		"params": [1, 2, 3]
	},
	{
		"jsonrpc":"2.0",
		"method":"sum",
		"id": 1,
		"params": [4, 5, 6]
	}]`))
	if !isBatch {
		t.Error("should be a batch")
		return
	}
	if len(responses) != 1 {
		t.Errorf("should be one response but got %v", len(responses))
		return
	}
	if invoked != 2 {
		t.Errorf("method not invoked enouhg times: %v but expected 2 times", invoked)
	}
	if bytes.Compare(responses[0].ID, []byte("1")) != 0 {
		t.Errorf("unmatched ID: got %v", string(responses[0].ID))
	}
	if responses[0].Version != "2.0" {
		t.Errorf("version is not 2.0: %v", responses[0].Version)
	}
	if responses[0].Error != nil {
		t.Errorf("failed: %+v", responses[0].Error)
		return
	}
	if v, ok := responses[0].Result.(int); !ok || v != 15 {
		t.Errorf("not matched result: %v", responses[0].Result)
	}
}
func TestNoMethod(t *testing.T) {
	router := &Router{}
	responses, isBatch := router.Invoke(bytes.NewBufferString(`{
		"jsonrpc":"2.0",
		"method":"sum",
		"id": 1,
		"params": [1, 2, 3]
	}`))
	if isBatch {
		t.Error("should be not a batch")
		return
	}
	if len(responses) != 1 {
		t.Errorf("only one response expected but got %v", len(responses))
		return
	}
	if bytes.Compare(responses[0].ID, []byte("1")) != 0 {
		t.Errorf("unmatched ID: got %v", string(responses[0].ID))
	}
	if responses[0].Version != "2.0" {
		t.Errorf("version is not 2.0: %v", responses[0].Version)
	}
	if responses[0].Error == nil {
		t.Errorf("should be failed but got response: %+v", responses[0].Result)
		return
	}
	if responses[0].Error.Code != MethodNotFound {
		t.Errorf("error code should be NotFound but got: %+v", responses[0].Error.Code)
		return
	}
}

func TestBrokenJSON(t *testing.T) {
	router := &Router{}
	responses, isBatch := router.Invoke(bytes.NewBufferString(`{
		"jsonrpc":"2.0",
		"method":"sum",
		"id": 1,
		"params": [1, 2, 3],
		trashData
	}`))
	if isBatch {
		t.Error("should be not a batch")
		return
	}
	if len(responses) != 1 {
		t.Errorf("only one response expected but got %v", len(responses))
		return
	}
	if responses[0].ID != nil {
		t.Errorf("should be no ID: got %v", string(responses[0].ID))
	}
	if responses[0].Version != "2.0" {
		t.Errorf("version is not 2.0: %v", responses[0].Version)
	}
	if responses[0].Error == nil {
		t.Errorf("should be failed but got response: %+v", responses[0].Result)
		return
	}
	if responses[0].Error.Code != ParseError {
		t.Errorf("error code should be ParseError but got: %+v", responses[0].Error.Code)
		return
	}
}

func TestInvalidRequest(t *testing.T) {
	router := &Router{}
	responses, isBatch := router.Invoke(bytes.NewBufferString(`{
		"jsonrpc":"1.0",
		"method":"sum",
		"id": 1,
		"params": [1, 2, 3]
	}`))
	if isBatch {
		t.Error("should be not a batch")
		return
	}
	if len(responses) != 1 {
		t.Errorf("only one response expected but got %v", len(responses))
		return
	}
	if bytes.Compare(responses[0].ID, []byte("1")) != 0 {
		t.Errorf("unmatched ID: got %v", string(responses[0].ID))
	}
	if responses[0].Version != "2.0" {
		t.Errorf("version is not 2.0: %v", responses[0].Version)
	}
	if responses[0].Error == nil {
		t.Errorf("should be failed but got response: %+v", responses[0].Result)
		return
	}
	if responses[0].Error.Code != InvalidRequest {
		t.Errorf("error code should be InvalidRequest but got: %+v", responses[0].Error.Code)
		return
	}
}

func TestAppError(t *testing.T) {
	router := &Router{}
	err := router.RegisterPositionalOnly("sum", func(a, b int, c *int) (int, error) {
		return a + b + *c, errors.New("the summation error")
	})
	if err != nil {
		t.Error("reg function", err)
		return
	}
	responses, isBatch := router.Invoke(bytes.NewBufferString(`{
		"jsonrpc":"2.0",
		"method":"sum",
		"id": 1,
		"params": [1, 2, 3]
	}`))
	if isBatch {
		t.Error("should be not a batch")
		return
	}
	if len(responses) != 1 {
		t.Errorf("only one response expected but got %v", len(responses))
		return
	}
	if bytes.Compare(responses[0].ID, []byte("1")) != 0 {
		t.Errorf("unmatched ID: got %v", string(responses[0].ID))
	}
	if responses[0].Version != "2.0" {
		t.Errorf("version is not 2.0: %v", responses[0].Version)
	}
	if responses[0].Error == nil {
		t.Errorf("should be failed but got response: %+v", responses[0].Result)
		return
	}
	if responses[0].Result != nil {
		t.Errorf("should be empty result but got: %+v", responses[0].Result)
	}
	if responses[0].Error.Code != AppError {
		t.Errorf("error code should be AppError but got: %+v", responses[0].Error.Code)
		return
	}
	if responses[0].Error.Message != "the summation error" {
		t.Errorf("error message unexpectable got: %+v", responses[0].Error.Message)
		return
	}
}

func TestHandler(t *testing.T) {
	router := &Router{}
	err := router.RegisterPositionalOnly("sum", func(a, b int, c *int) (int, error) {
		return a + b + *c, nil
	})
	if err != nil {
		t.Error("reg function", err)
		return
	}
	server := httptest.NewServer(Handler(router))
	defer server.Close()

	res, err := http.Post(server.URL, "application/json", bytes.NewBufferString(`{
		"jsonrpc":"2.0",
		"id": 1,
		"method":"sum",
		"params": [1, 2, 3]
	}`))
	if err != nil {
		t.Errorf("request over http: %v", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("status: %v", res.StatusCode)
		return
	}
	var response Response
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		t.Errorf("unmarshal: %v", err)
		return
	}

	if bytes.Compare(response.ID, []byte("1")) != 0 {
		t.Errorf("unmatched ID: got %v", string(response.ID))
	}
	if response.Version != "2.0" {
		t.Errorf("version is not 2.0: %v", response.Version)
	}
	if response.Error != nil {
		t.Errorf("failed: %+v", response.Error)
		return
	}
	if v, ok := response.Result.(float64); !ok || v != 6 {
		t.Errorf("not matched result: %v (%v)", response.Result, reflect.TypeOf(response.Result).Name())
	}
}

func TestRouter_InterceptMethods(t *testing.T) {
	router := &Router{}
	err := router.RegisterPositionalOnly("sum", func(a, b int, c *int) (int, error) {
		return a + b + *c, nil
	})
	if err != nil {
		t.Error("reg function", err)
		return
	}
	var complete bool
	router.InterceptMethods(func(ic *MethodInterceptorContext) (i interface{}, err error) {
		a, b := ic.Next()
		complete = true
		return a, b
	})
	responses, isBatch := router.Invoke(bytes.NewBufferString(`{
		"jsonrpc":"2.0",
		"id": 1,
		"method":"sum",
		"params": [1, 2, 3]
	}`))
	if isBatch {
		t.Error("should be not a batch")
		return
	}
	if len(responses) != 1 {
		t.Errorf("only one response expected but got %v", len(responses))
		return
	}
	if bytes.Compare(responses[0].ID, []byte("1")) != 0 {
		t.Errorf("unmatched ID: got %v", string(responses[0].ID))
	}
	if responses[0].Version != "2.0" {
		t.Errorf("version is not 2.0: %v", responses[0].Version)
	}
	if responses[0].Error != nil {
		t.Errorf("failed: %+v", responses[0].Error)
		return
	}
	if v, ok := responses[0].Result.(int); !ok || v != 6 {
		t.Errorf("not matched result: %v", responses[0].Result)
	}
	if !complete {
		t.Error("interceptor not invoked")
	}
}
func TestRouter_Intercept(t *testing.T) {
	router := &Router{}
	err := router.RegisterPositionalOnly("sum", func(a, b int, c *int) (int, error) {
		return a + b + *c, nil
	})
	if err != nil {
		t.Error("reg function", err)
		return
	}
	var complete bool
	router.Intercept(func(gic *GlobalInterceptorContext) (responses []*Response, isBatch bool) {
		complete = len(gic.Requests) == 1 && !gic.IsBatch
		return gic.Next()
	})
	responses, isBatch := router.Invoke(bytes.NewBufferString(`{
		"jsonrpc":"2.0",
		"id": 1,
		"method":"sum",
		"params": [1, 2, 3]
	}`))
	if isBatch {
		t.Error("should be not a batch")
		return
	}
	if len(responses) != 1 {
		t.Errorf("only one response expected but got %v", len(responses))
		return
	}
	if bytes.Compare(responses[0].ID, []byte("1")) != 0 {
		t.Errorf("unmatched ID: got %v", string(responses[0].ID))
	}
	if responses[0].Version != "2.0" {
		t.Errorf("version is not 2.0: %v", responses[0].Version)
	}
	if responses[0].Error != nil {
		t.Errorf("failed: %+v", responses[0].Error)
		return
	}
	if v, ok := responses[0].Result.(int); !ok || v != 6 {
		t.Errorf("not matched result: %v", responses[0].Result)
	}
	if !complete {
		t.Error("interceptor not invoked")
	}
}
func TestCustomError(t *testing.T) {
	router := &Router{}
	err := router.RegisterPositionalOnly("sum", func(a, b int, c *int) (int, error) {
		return a + b + *c, &Error{
			Code:    1234,
			Message: "Test Error",
		}
	})
	if err != nil {
		t.Error("reg function", err)
		return
	}
	responses, isBatch := router.Invoke(bytes.NewBufferString(`{
		"jsonrpc":"2.0",
		"id": 1,
		"method":"sum",
		"params": [1, 2, 3]
	}`))
	if isBatch {
		t.Error("should be not a batch")
		return
	}
	if len(responses) != 1 {
		t.Errorf("only one response expected but got %v", len(responses))
		return
	}
	if bytes.Compare(responses[0].ID, []byte("1")) != 0 {
		t.Errorf("unmatched ID: got %v", string(responses[0].ID))
	}
	if responses[0].Version != "2.0" {
		t.Errorf("version is not 2.0: %v", responses[0].Version)
	}
	if responses[0].Error == nil {
		t.Errorf("should be failed but got response: %+v", responses[0].Result)
		return
	}
	if responses[0].Result != nil {
		t.Errorf("should be empty result but got: %+v", responses[0].Result)
	}
	if responses[0].Error.Code != 1234 {
		t.Errorf("error code should be 1234 but got: %+v", responses[0].Error.Code)
		return
	}
	if responses[0].Error.Message != "Test Error" {
		t.Errorf("error message unexpectable got: %+v", responses[0].Error.Message)
		return
	}
}

func ExampleRouter_RegisterPositionalOnly() {
	router := &Router{}
	err := router.RegisterPositionalOnly("sum", func(a, b int) (int, error) {
		return a + b, nil
	})
	if err != nil {
		panic(err)
	}
}

func ExampleRouter_RegisterNamedOnly() {
	type Args struct {
		A int `json:"a"`
		B int `json:"b"`
	}
	router := &Router{}
	err := router.RegisterNamedOnly("sum", func(params *Args) (int, error) {
		return params.A + params.B, nil
	})
	if err != nil {
		panic(err)
	}
}

func ExampleHandler() {
	var router Router
	router.RegisterPositionalOnly("sum", func(a, b int) (int, error) {
		return a + b, nil
	})
	http.ListenAndServe(":8080", Handler(&router))
}
