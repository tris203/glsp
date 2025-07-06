package glsp

import (
	"context"
	"encoding/json"
	"testing"
)

func TestLspProtocolVersionConstants(t *testing.T) {
	if Protocol_316 != 0 {
		t.Errorf("Protocol_316 should be 0, got %d", Protocol_316)
	}
	if Protocol_317 != 1 {
		t.Errorf("Protocol_317 should be 1, got %d", Protocol_317)
	}
	if Protocol_318 != 2 {
		t.Errorf("Protocol_318 should be 2, got %d", Protocol_318)
	}
}

func TestContext(t *testing.T) {
	params := json.RawMessage(`{"test": "value"}`)
	ctx := Context{
		Method:           "test/method",
		Params:           params,
		Context:          context.Background(),
		Protocol_Version: Protocol_316,
	}

	if ctx.Method != "test/method" {
		t.Errorf("Expected method 'test/method', got '%s'", ctx.Method)
	}
	if ctx.Protocol_Version != Protocol_316 {
		t.Errorf("Expected protocol version %d, got %d", Protocol_316, ctx.Protocol_Version)
	}
	if ctx.Context == nil {
		t.Error("Context should not be nil")
	}
	if string(ctx.Params) != `{"test": "value"}` {
		t.Errorf("Expected params to be preserved, got %s", string(ctx.Params))
	}
}

func TestContextWithNotifyAndCall(t *testing.T) {
	notifyCalled := false
	callCalled := false

	ctx := Context{
		Method: "test/method",
		Notify: func(method string, params any) {
			notifyCalled = true
			if method != "test/notify" {
				t.Errorf("Expected notify method 'test/notify', got '%s'", method)
			}
		},
		Call: func(method string, params any, result any) {
			callCalled = true
			if method != "test/call" {
				t.Errorf("Expected call method 'test/call', got '%s'", method)
			}
		},
	}

	if ctx.Method != "test/method" {
		t.Errorf("Expected method 'test/method', got '%s'", ctx.Method)
	}

	ctx.Notify("test/notify", nil)
	ctx.Call("test/call", nil, nil)

	if !notifyCalled {
		t.Error("Notify function was not called")
	}
	if !callCalled {
		t.Error("Call function was not called")
	}
}
