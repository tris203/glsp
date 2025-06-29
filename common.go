package glsp

import (
	contextpkg "context"
	"encoding/json"
)

type LspProtocolVersion int

const (
	Protocol_3_16 LspProtocolVersion = iota
	Protocol_3_17
	Protocol_3_18
)

type NotifyFunc func(method string, params any)
type CallFunc func(method string, params any, result any)

type Context struct {
	Method           string
	Params           json.RawMessage
	Notify           NotifyFunc
	Call             CallFunc
	Context          contextpkg.Context // can be nil
	Protocol_Version LspProtocolVersion
}

type RPCHandler interface {
	IsInitialized() bool
	SetInitialized(initialized bool)
	GetVersion() LspProtocolVersion
	GetMethodMap() map[string]HandlerInterface
	GetCustomMethods() map[string]HandlerInterface
}

type HandlerInterface interface {
	Handle(*Context, []byte) (any, error)
}
