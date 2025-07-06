//go:generate go run ../tools/genprotocol/main.go
package protocol

import (
	"fmt"
	"sync"

	"github.com/tliron/glsp"
	"github.com/tliron/glsp/protocol/translation"
)

var customMethods = make(map[string]glsp.HandlerInterface)
var mu sync.RWMutex

type CustomMethodProvider interface {
	GetCustomMethods() map[string]glsp.HandlerInterface
}

func (h *Handler_316) GetCustomMethods() map[string]glsp.HandlerInterface {
	mu.RLock()
	defer mu.RUnlock()
	return customMethods
}

func (h *Handler_317) GetCustomMethods() map[string]glsp.HandlerInterface {
	mu.RLock()
	defer mu.RUnlock()
	return customMethods
}

func AddCustomRequest[P any, R any](method string, handler RequestFunc[P, R]) error {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := customMethods[method]; exists {
		return fmt.Errorf("method %q already registered", method)
	}
	customMethods[method] = translation.NewTypedHandler(handler)
	return nil
}

func AddCustomNotification[P any](method string, handler NotificationFunc[P]) error {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := customMethods[method]; exists {
		return fmt.Errorf("method %q already registered", method)
	}
	customMethods[method] = translation.NewNotificationHandler(handler)
	return nil
}
