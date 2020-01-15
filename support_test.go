package jsonrpc2

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestMaxBatch(t *testing.T) {
	router := &Router{}
	err := router.RegisterPositionalOnly("sum", func(a, b int, c *int) (int, error) {
		return a + b + *c, nil
	})
	if err != nil {
		t.Error("reg function", err)
		return
	}
	router.Intercept(MaxBatch(2))
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
	},
	{
		"jsonrpc":"2.0",
		"id": 3,
		"method":"sum",
		"params": [7, 8, 9]
	}]`))
	if !isBatch {
		t.Error("should be  batch")
		return
	}
	if len(responses) != 1 {
		t.Errorf("only one response expected but got %v", len(responses))
		return
	}
	if responses[0].Version != "2.0" {
		t.Errorf("version is not 2.0: %v", responses[1].Version)
	}
	if responses[0].Error == nil {
		t.Errorf("(0) not failed")
		return
	}
	if responses[0].Error.Code != InternalError {
		t.Errorf("error code should be InternalError but got %d", responses[0].Error.Code)
		return
	}
	if responses[0].Error.Message != "batch is too big" {
		t.Errorf("error message unexpected: got %s", responses[0].Error.Message)
		return
	}
}

func TestToArray(t *testing.T) {
	var raw json.RawMessage
	err := json.Unmarshal([]byte("[1, 2, 3]"), &raw)
	if err != nil {
		t.Fatal(err)
	}
	array, err := ToArray(raw, 2)
	if err == nil {
		t.Error("should return error if len not matched")
		return
	}
	array, err = ToArray(raw, 3)
	if err != nil {
		t.Errorf("unmatched: %v", err)
		return
	}

	if len(array) != 3 {
		t.Errorf("array should 3 but got %v", len(array))
	}
}

func TestUnmarshalArray(t *testing.T) {
	var raw json.RawMessage
	err := json.Unmarshal([]byte("[1, 2, 3]"), &raw)
	if err != nil {
		t.Fatal(err)
	}
	var a, b, c int
	err = UnmarshalArray(raw, &a, &b, &c)
	if err != nil {
		t.Errorf("unmatched: %v", err)
		return
	}

	if a != 1 || b != 2 || c != 3 {
		t.Errorf("got %v, %v, %v", a, b, c)
	}
}
