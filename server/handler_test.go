package server

import (
	"encoding/json"
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/tliron/glsp"
	"github.com/tliron/glsp/protocol"
)

type mockHandlerInterface struct {
	handleFunc func(*glsp.Context, []byte) (any, error)
}

func (m *mockHandlerInterface) Handle(ctx *glsp.Context, params []byte) (any, error) {
	if m.handleFunc != nil {
		return m.handleFunc(ctx, params)
	}
	return "mock response", nil
}

func TestServer_HandlerHandle_NotInitialized(t *testing.T) {
	handler := &mockHandler{
		initialized:   false,
		version:       glsp.Protocol_316,
		methodMap:     make(map[string]glsp.HandlerInterface),
		customMethods: make(map[string]glsp.HandlerInterface),
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	server := NewServer(handler, logger, false)

	ctx := &glsp.Context{
		Method: "textDocument/hover",
	}

	_, validMethod, validParams, err := server.HandlerHandle(ctx)
	if err == nil {
		t.Error("Expected error for uninitialized server, got nil")
	}
	if err.Error() != "server not initialized" {
		t.Errorf("Expected 'server not initialized' error, got '%s'", err.Error())
	}
	if !validMethod {
		t.Error("Expected validMethod to be true")
	}
	if !validParams {
		t.Error("Expected validParams to be true")
	}
}

func TestServer_HandlerHandle_InitializeMethod(t *testing.T) {
	mockHandlerIface := &mockHandlerInterface{
		handleFunc: func(ctx *glsp.Context, params []byte) (any, error) {
			return "initialize response", nil
		},
	}

	handler := &mockHandler{
		initialized: false,
		version:     glsp.Protocol_316,
		methodMap: map[string]glsp.HandlerInterface{
			protocol.MethodInitialize: mockHandlerIface,
		},
		customMethods: make(map[string]glsp.HandlerInterface),
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	server := NewServer(handler, logger, false)

	ctx := &glsp.Context{
		Method: protocol.MethodInitialize,
		Params: json.RawMessage(`{}`),
	}

	result, validMethod, validParams, err := server.HandlerHandle(ctx)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !validMethod {
		t.Error("Expected validMethod to be true")
	}
	if !validParams {
		t.Error("Expected validParams to be true")
	}
	if result != "initialize response" {
		t.Errorf("Expected 'initialize response', got '%v'", result)
	}
	if !handler.IsInitialized() {
		t.Error("Handler should be initialized after successful Initialize method")
	}
}

func TestServer_HandlerHandle_MethodNotFound(t *testing.T) {
	handler := &mockHandler{
		initialized:   true,
		version:       glsp.Protocol_316,
		methodMap:     make(map[string]glsp.HandlerInterface),
		customMethods: make(map[string]glsp.HandlerInterface),
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	server := NewServer(handler, logger, false)

	ctx := &glsp.Context{
		Method: "nonexistent/method",
	}

	_, validMethod, validParams, err := server.HandlerHandle(ctx)
	if err == nil {
		t.Error("Expected error for nonexistent method, got nil")
	}
	if validMethod {
		t.Error("Expected validMethod to be false for nonexistent method")
	}
	if validParams {
		t.Error("Expected validParams to be false for nonexistent method")
	}
}

func TestServer_HandlerHandle_CustomMethod(t *testing.T) {

	type customParams struct {
		This string `json:"this"`
	}
	mockHandlerIface := &mockHandlerInterface{
		handleFunc: func(ctx *glsp.Context, params []byte) (any, error) {
			unmarshaledParams := &customParams{}
			if err := json.Unmarshal(params, unmarshaledParams); err != nil {
				return nil, err
			}
			if unmarshaledParams.This != "that" {
				t.Errorf("Expected 'this' to be 'that', got '%s'", unmarshaledParams.This)
			}

			return "custom response", nil
		},
	}

	handler := &mockHandler{
		initialized: true,
		version:     glsp.Protocol_316,
		methodMap:   make(map[string]glsp.HandlerInterface),
		customMethods: map[string]glsp.HandlerInterface{
			"custom/method": mockHandlerIface,
		},
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	server := NewServer(handler, logger, false)

	ctx := &glsp.Context{
		Method: "custom/method",
		Params: json.RawMessage(`{"this": "that"}`),
	}

	result, validMethod, validParams, err := server.HandlerHandle(ctx)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !validMethod {
		t.Error("Expected validMethod to be true for custom method")
	}
	if !validParams {
		t.Error("Expected validParams to be true for custom method")
	}
	if result != "custom response" {
		t.Errorf("Expected 'custom response', got '%v'", result)
	}
}

func TestServer_HandlerHandle_CustomMethodsAreHandlerScoped(t *testing.T) {
	type customParams struct {
		Value string `json:"value"`
	}

	handlerA := &protocol.Handler_316{}
	handlerA.SetInitialized(true)
	if err := protocol.AddCustomRequest(handlerA, "custom/scoped", func(ctx *glsp.Context, params *customParams) (string, error) {
		return "handler-a:" + params.Value, nil
	}); err != nil {
		t.Fatalf("Unexpected handler A registration error: %v", err)
	}

	handlerB := &protocol.Handler_316{}
	handlerB.SetInitialized(true)
	if err := protocol.AddCustomRequest(handlerB, "custom/scoped", func(ctx *glsp.Context, params *customParams) (string, error) {
		return "handler-b:" + params.Value, nil
	}); err != nil {
		t.Fatalf("Unexpected handler B registration error: %v", err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	serverA := NewServer(handlerA, logger, false)
	serverB := NewServer(handlerB, logger, false)
	ctx := &glsp.Context{
		Method: "custom/scoped",
		Params: json.RawMessage(`{"value":"ok"}`),
	}

	resultA, validMethod, validParams, err := serverA.HandlerHandle(ctx)
	if err != nil {
		t.Fatalf("Unexpected handler A error: %v", err)
	}
	if !validMethod || !validParams {
		t.Fatalf("Expected handler A custom method to be valid, got validMethod=%v validParams=%v", validMethod, validParams)
	}
	if resultA != "handler-a:ok" {
		t.Errorf("Expected handler A result, got %v", resultA)
	}

	resultB, validMethod, validParams, err := serverB.HandlerHandle(ctx)
	if err != nil {
		t.Fatalf("Unexpected handler B error: %v", err)
	}
	if !validMethod || !validParams {
		t.Fatalf("Expected handler B custom method to be valid, got validMethod=%v validParams=%v", validMethod, validParams)
	}
	if resultB != "handler-b:ok" {
		t.Errorf("Expected handler B result, got %v", resultB)
	}
}

func TestServer_HandlerHandle_CustomMethodDoesNotLeak(t *testing.T) {
	type customParams struct{}

	handlerA := &protocol.Handler_316{}
	handlerA.SetInitialized(true)
	if err := protocol.AddCustomRequest(handlerA, "custom/private", func(ctx *glsp.Context, params *customParams) (string, error) {
		return "private", nil
	}); err != nil {
		t.Fatalf("Unexpected registration error: %v", err)
	}

	handlerB := &protocol.Handler_316{}
	handlerB.SetInitialized(true)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	serverB := NewServer(handlerB, logger, false)

	_, validMethod, validParams, err := serverB.HandlerHandle(&glsp.Context{Method: "custom/private", Params: json.RawMessage(`{}`)})
	if err == nil {
		t.Fatal("Expected unregistered custom method to fail")
	}
	if validMethod {
		t.Error("Expected custom method not to leak to handler B")
	}
	if validParams {
		t.Error("Expected invalid params when method is not found")
	}
}

func TestServer_HandlerHandle_HandlerError(t *testing.T) {
	mockHandlerIface := &mockHandlerInterface{
		handleFunc: func(ctx *glsp.Context, params []byte) (any, error) {
			return nil, errors.New("handler error")
		},
	}

	handler := &mockHandler{
		initialized: true,
		version:     glsp.Protocol_316,
		methodMap: map[string]glsp.HandlerInterface{
			"test/method": mockHandlerIface,
		},
		customMethods: make(map[string]glsp.HandlerInterface),
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	server := NewServer(handler, logger, false)

	ctx := &glsp.Context{
		Method: "test/method",
		Params: json.RawMessage(`{}`),
	}

	_, validMethod, validParams, err := server.HandlerHandle(ctx)
	if err == nil {
		t.Error("Expected error from handler, got nil")
	}
	if err.Error() != "handler error" {
		t.Errorf("Expected 'handler error', got '%s'", err.Error())
	}
	if !validMethod {
		t.Error("Expected validMethod to be true")
	}
	if !validParams {
		t.Error("Expected validParams to be true")
	}
}

func TestServer_HandlerHandle_UnmarshalError(t *testing.T) {
	mockHandlerIface := &mockHandlerInterface{
		handleFunc: func(ctx *glsp.Context, params []byte) (any, error) {
			return nil, &json.UnmarshalTypeError{
				Value: "string",
				Type:  nil,
			}
		},
	}

	handler := &mockHandler{
		initialized: true,
		version:     glsp.Protocol_316,
		methodMap: map[string]glsp.HandlerInterface{
			"test/method": mockHandlerIface,
		},
		customMethods: make(map[string]glsp.HandlerInterface),
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	server := NewServer(handler, logger, false)

	ctx := &glsp.Context{
		Method: "test/method",
		Params: json.RawMessage(`{}`),
	}

	_, validMethod, validParams, err := server.HandlerHandle(ctx)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !validMethod {
		t.Error("Expected validMethod to be true")
	}
	if validParams {
		t.Error("Expected validParams to be false for unmarshal error")
	}
}

func TestServer_HandlerHandle_SuccessfulMethod(t *testing.T) {
	mockHandlerIface := &mockHandlerInterface{
		handleFunc: func(ctx *glsp.Context, params []byte) (any, error) {
			return "success response", nil
		},
	}

	handler := &mockHandler{
		initialized: true,
		version:     glsp.Protocol_316,
		methodMap: map[string]glsp.HandlerInterface{
			"test/method": mockHandlerIface,
		},
		customMethods: make(map[string]glsp.HandlerInterface),
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	server := NewServer(handler, logger, false)

	ctx := &glsp.Context{
		Method: "test/method",
		Params: json.RawMessage(`{}`),
	}

	result, validMethod, validParams, err := server.HandlerHandle(ctx)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !validMethod {
		t.Error("Expected validMethod to be true")
	}
	if !validParams {
		t.Error("Expected validParams to be true")
	}
	if result != "success response" {
		t.Errorf("Expected 'success response', got '%v'", result)
	}
}

func TestServer_HandlerHandle_ShutdownWithoutParams(t *testing.T) {
	called := false
	handler := &protocol.Handler_316{
		Shutdown: func(ctx *glsp.Context) error {
			called = true
			return nil
		},
	}
	handler.SetInitialized(true)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	server := NewServer(handler, logger, false)

	_, validMethod, validParams, err := server.HandlerHandle(&glsp.Context{Method: protocol.MethodShutdown})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !validMethod {
		t.Error("Expected shutdown to be a valid method")
	}
	if !validParams {
		t.Error("Expected shutdown without params to have valid params")
	}
	if !called {
		t.Error("Expected shutdown handler to be called")
	}
	if handler.IsInitialized() {
		t.Error("Expected shutdown to clear initialized state")
	}
}

func TestServer_HandlerHandle_CodeActionUsesCodeActionParams(t *testing.T) {
	called := false
	handler := &protocol.Handler_316{
		TextDocumentCodeAction: func(ctx *glsp.Context, params *protocol.CodeActionParams) ([]protocol.CodeAction, error) {
			called = true
			if params.TextDocument.URI != "file:///test.go" {
				t.Errorf("Expected text document URI to unmarshal, got %q", params.TextDocument.URI)
			}
			return []protocol.CodeAction{}, nil
		},
	}
	handler.SetInitialized(true)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	server := NewServer(handler, logger, false)

	params := json.RawMessage(`{"textDocument":{"uri":"file:///test.go"},"range":{"start":{"line":1,"character":2},"end":{"line":1,"character":3}},"context":{"diagnostics":[]}}`)
	_, validMethod, validParams, err := server.HandlerHandle(&glsp.Context{Method: protocol.MethodTextDocumentCodeAction, Params: params})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !validMethod || !validParams {
		t.Fatalf("Expected valid method and params, got validMethod=%v validParams=%v", validMethod, validParams)
	}
	if !called {
		t.Error("Expected code action handler to be called")
	}
}
