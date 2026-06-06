//go:generate go run ../tools/gopls_generator

package protocol

import (
	"github.com/tliron/glsp"
	"sync"
)

type common_handler struct {
	initialized   bool
	lock          sync.Mutex
	handlerMap    map[LSPMethod]glsp.HandlerInterface
	customMethods map[string]glsp.HandlerInterface
}

type RequestFunc[P any, R any] func(context *glsp.Context, params *P) (R, error)
type NotificationFunc[P any] func(context *glsp.Context, params *P) error
type ContextOnlyFunc func(context *glsp.Context) error
