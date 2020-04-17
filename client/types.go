package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/reddec/jsonrpc2"
	"net/http"
)

// Standard JSON-RPC 2.0 request messages
type request struct {
	// always 2.0 (will refuse if not)
	Version string `json:"jsonrpc"`
	// case-sensitive method name
	Method string `json:"method"`
	// any kind of valid JSON as ID (more relaxed comparing to for spec)
	ID interface{} `json:"id"`
	// array (for positional) or object (for named) of arguments
	Params []interface{} `json:"params"`
}

// Call JSON-RPC 2.0 over HTTP
func CallHTTP(ctx context.Context, url string, method string, id interface{}, out interface{}, params ...interface{}) error {
	data, err := json.Marshal(&request{
		Version: "2.0",
		Method:  method,
		ID:      id,
		Params:  params,
	})
	if err != nil {
		return fmt.Errorf("%s [id=%v]: call over HTTP via %s, JSON encode: %w", method, id, url, err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("%s [id=%v]: call over HTTP via %s, preapre request: %w", method, id, url, err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("%s [id=%v]: call over HTTP via %s, make request: %w", method, id, url, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s [id=%v]: call over HTTP via %s, request status code %d - %s", method, id, url, res.StatusCode, res.Status)
	}
	var reply jsonrpc2.Response
	reply.Result = out
	err = json.NewDecoder(res.Body).Decode(&reply)
	if err != nil {
		return fmt.Errorf("%s [id=%v]: call over HTTP via %s, parse response: %w", method, id, url, err)
	}
	if reply.Error != nil {
		return reply.Error
	}
	return nil
}
