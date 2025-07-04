package translation

import (
	"encoding/json"
	"github.com/tliron/glsp"
)

type TranslationFunc[P, R any] func(*glsp.Context, *P) (R, error)

// Implement the HandlerInterface
func (h TranslationFunc[P, R]) Handle(ctx *glsp.Context, params []byte) (any, error) {
	var p P
	if err := json.Unmarshal(params, &p); err != nil {
		return nil, err
	}
	result, err := h(ctx, &p)
	if err != nil {
		return nil, err
	}

	// Handle the case where R is struct{} (void return)
	var zero R
	if any(zero) == any(struct{}{}) {
		return nil, nil
	}

	return result, nil
}

// Factory function for notification handlers (return only error)
func NewNotificationHandler[P any](fn func(*glsp.Context, *P) error) glsp.HandlerInterface {
	return TranslationFunc[P, struct{}](func(ctx *glsp.Context, p *P) (struct{}, error) {
		return struct{}{}, fn(ctx, p)
	})
}

// Factory function for context-only handlers (no parameters)
func NewContextOnlyHandler(fn func(*glsp.Context) error) glsp.HandlerInterface {
	return TranslationFunc[struct{}, struct{}](func(ctx *glsp.Context, _ *struct{}) (struct{}, error) {
		return struct{}{}, fn(ctx)
	})
}

// Factory function for handlers with specific return types - converts to any
func NewTypedHandler[P, R any](fn func(*glsp.Context, *P) (R, error)) glsp.HandlerInterface {
	return TranslationFunc[P, any](func(ctx *glsp.Context, p *P) (any, error) {
		result, err := fn(ctx, p)
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

// Factory function for handlers with non-specific return types - converts to any
func NewCustomHandler(fn func(ctx *glsp.Context, p any) (any, error)) glsp.HandlerInterface {
	return TranslationFunc[any, any](func(ctx *glsp.Context, p *any) (any, error) {
		result, err := fn(ctx, p)
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}
