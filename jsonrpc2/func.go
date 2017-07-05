package jsonrpc2

import (
	"context"
	"encoding/json"
)

// Func links a method of JSON-RPC request.
type Func func(c context.Context, params *json.RawMessage) (result interface{}, err *Error)
