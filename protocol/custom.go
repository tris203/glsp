package protocol

import (
	"encoding/json"

	"github.com/tliron/glsp"
)

type CustomMethodHandler struct {
	Func   CustomMethodFunc
	params json.RawMessage
}

type CustomMethodHandlers map[string]CustomMethodHandler

type CustomMethodFunc func(context *glsp.Context, params json.RawMessage) (any, error)
