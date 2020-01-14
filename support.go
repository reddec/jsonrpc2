package jsonrpc2

import (
	"encoding/json"
	"fmt"
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
