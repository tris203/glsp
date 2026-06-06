//go:generate go run ../tools/genprotocol/main.go
package protocol

import (
	"fmt"

	"github.com/tliron/glsp"
	"github.com/tliron/glsp/protocol/translation"
)

type CustomMethodProvider interface {
	GetCustomMethods() map[string]glsp.HandlerInterface
	AddCustomMethod(method string, handler glsp.HandlerInterface) error
}

func (h *common_handler) GetCustomMethods() map[string]glsp.HandlerInterface {
	h.lock.Lock()
	defer h.lock.Unlock()

	methods := make(map[string]glsp.HandlerInterface, len(h.customMethods))
	for method, handler := range h.customMethods {
		methods[method] = handler
	}
	return methods
}

func (h *common_handler) AddCustomMethod(method string, handler glsp.HandlerInterface) error {
	h.lock.Lock()
	defer h.lock.Unlock()

	if h.customMethods == nil {
		h.customMethods = make(map[string]glsp.HandlerInterface)
	}
	if _, exists := h.customMethods[method]; exists {
		return fmt.Errorf("method %q already registered", method)
	}
	h.customMethods[method] = handler
	return nil
}

func AddCustomRequest[P any, R any](provider CustomMethodProvider, method string, handler RequestFunc[P, R]) error {
	return provider.AddCustomMethod(method, translation.NewTypedHandler(handler))
}

func AddCustomNotification[P any](provider CustomMethodProvider, method string, handler NotificationFunc[P]) error {
	return provider.AddCustomMethod(method, translation.NewNotificationHandler(handler))
}
