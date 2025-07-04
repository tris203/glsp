package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sourcegraph/jsonrpc2"

	"github.com/tliron/glsp"
	"github.com/tliron/glsp/protocol"
)

// See: https://github.com/sourcegraph/go-langserver/blob/master/langserver/handler.go#L206

func (s *Server) newHandler() jsonrpc2.Handler {
	return jsonrpc2.HandlerWithError(s.handle)
}

func (s *Server) handle(ctx context.Context, connection *jsonrpc2.Conn, request *jsonrpc2.Request) (any, error) {
	glspContext := glsp.Context{
		Method: request.Method,
		Notify: func(method string, params any) {
			if err := connection.Notify(ctx, method, params); err != nil {
				s.Log.Error(err.Error())
			}
		},
		Call: func(method string, params any, result any) {
			if err := connection.Call(ctx, method, params, result); err != nil {
				s.Log.Error(err.Error())
			}
		},
		Context:          ctx,
		Protocol_Version: s.Handler.GetVersion(),
	}

	if request.Params != nil {
		glspContext.Params = *request.Params
	}

	switch request.Method {
	case "exit":
		// We're giving the attached handler a chance to handle it first, but we'll ignore any result
		s.HandlerHandle(&glspContext)
		err := connection.Close()
		return nil, err

	default:
		// Note: jsonrpc2 will not even call this function if reqest.Params is invalid JSON,
		// so we don't need to handle jsonrpc2.CodeParseError here
		result, validMethod, validParams, err := s.HandlerHandle(&glspContext)
		if !validMethod {
			return nil, &jsonrpc2.Error{
				Code:    jsonrpc2.CodeMethodNotFound,
				Message: fmt.Sprintf("method not supported: %s", request.Method),
			}
		} else if !validParams {
			if err == nil {
				return nil, &jsonrpc2.Error{
					Code: jsonrpc2.CodeInvalidParams,
				}
			} else {
				return nil, &jsonrpc2.Error{
					Code:    jsonrpc2.CodeInvalidParams,
					Message: err.Error(),
				}
			}
		} else if err != nil {
			return nil, &jsonrpc2.Error{
				Code:    jsonrpc2.CodeInvalidRequest,
				Message: err.Error(),
			}
		} else {
			return result, nil
		}
	}
}

// ([glsp.Handler] interface)
func (s *Server) HandlerHandle(context *glsp.Context) (r any, validMethod bool, validParams bool, err error) {
	if !s.Handler.IsInitialized() && (context.Method != protocol.MethodInitialize) {
		return nil, true, true, errors.New("server not initialized")
	}

	handler, exists := s.Handler.GetMethodMap()[context.Method]
	if !exists {
		// Check custom methods
		customMethods := s.Handler.GetCustomMethods()
		if len(customMethods) > 0 {
			if customHandler, ok := customMethods[context.Method]; ok {
				validMethod = true
				r, err = customHandler.Handle(context, context.Params)
				validParams = true
				return r, validMethod, validParams, err
			}
		}
		return nil, false, false, fmt.Errorf("method not supported: %s", context.Method)
	}

	validMethod = true
	r, err = handler.Handle(context, context.Params)
	if err != nil {
		// Check if it's an unmarshal error (invalid params) vs handler error
		var unmarshalErr *json.UnmarshalTypeError
		if errors.As(err, &unmarshalErr) {
			return nil, true, false, nil
		}
	}
	validParams = true

	// Special handling for Initialize method
	if context.Method == protocol.MethodInitialize && err == nil {
		s.Handler.SetInitialized(true)
	}

	return r, validMethod, validParams, err

}
