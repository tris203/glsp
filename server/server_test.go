package server

import (
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/tliron/glsp"
)

type mockHandler struct {
	initialized   bool
	version       glsp.LspProtocolVersion
	methodMap     map[string]glsp.HandlerInterface
	customMethods map[string]glsp.HandlerInterface
}

func (m *mockHandler) IsInitialized() bool {
	return m.initialized
}

func (m *mockHandler) SetInitialized(initialized bool) {
	m.initialized = initialized
}

func (m *mockHandler) GetVersion() glsp.LspProtocolVersion {
	return m.version
}

func (m *mockHandler) GetMethodMap() map[string]glsp.HandlerInterface {
	return m.methodMap
}

func (m *mockHandler) GetCustomMethods() map[string]glsp.HandlerInterface {
	return m.customMethods
}

func TestNewServer(t *testing.T) {
	handler := &mockHandler{
		version:       glsp.Protocol_316,
		methodMap:     make(map[string]glsp.HandlerInterface),
		customMethods: make(map[string]glsp.HandlerInterface),
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	server := NewServer(handler, logger, true)

	if server.Handler != handler {
		t.Error("Handler not set correctly")
	}
	if server.Log != logger {
		t.Error("Logger not set correctly")
	}
	if !server.Debug {
		t.Error("Debug flag not set correctly")
	}
	if server.Timeout != DefaultTimeout {
		t.Errorf("Expected timeout %v, got %v", DefaultTimeout, server.Timeout)
	}
	if server.ReadTimeout != DefaultTimeout {
		t.Errorf("Expected read timeout %v, got %v", DefaultTimeout, server.ReadTimeout)
	}
	if server.WriteTimeout != DefaultTimeout {
		t.Errorf("Expected write timeout %v, got %v", DefaultTimeout, server.WriteTimeout)
	}
	if server.StreamTimeout != DefaultTimeout {
		t.Errorf("Expected stream timeout %v, got %v", DefaultTimeout, server.StreamTimeout)
	}
	if server.WebSocketTimeout != DefaultTimeout {
		t.Errorf("Expected websocket timeout %v, got %v", DefaultTimeout, server.WebSocketTimeout)
	}
}

func TestServerDefaults(t *testing.T) {
	if DefaultTimeout != time.Minute {
		t.Errorf("Expected default timeout to be 1 minute, got %v", DefaultTimeout)
	}
}

func TestServerTimeoutConfiguration(t *testing.T) {
	handler := &mockHandler{}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	server := NewServer(handler, logger, false)

	customTimeout := 30 * time.Second
	server.Timeout = customTimeout
	server.ReadTimeout = customTimeout
	server.WriteTimeout = customTimeout
	server.StreamTimeout = customTimeout
	server.WebSocketTimeout = customTimeout

	if server.Timeout != customTimeout {
		t.Errorf("Expected timeout %v, got %v", customTimeout, server.Timeout)
	}
	if server.ReadTimeout != customTimeout {
		t.Errorf("Expected read timeout %v, got %v", customTimeout, server.ReadTimeout)
	}
	if server.WriteTimeout != customTimeout {
		t.Errorf("Expected write timeout %v, got %v", customTimeout, server.WriteTimeout)
	}
	if server.StreamTimeout != customTimeout {
		t.Errorf("Expected stream timeout %v, got %v", customTimeout, server.StreamTimeout)
	}
	if server.WebSocketTimeout != customTimeout {
		t.Errorf("Expected websocket timeout %v, got %v", customTimeout, server.WebSocketTimeout)
	}
}
