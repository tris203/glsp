package protocol

import (
	"encoding/json"
	"testing"
)

func TestInitializeParams_316(t *testing.T) {
	processID := Integer(1234)
	params := InitializeParams_316{
		ProcessID: &processID,
		ClientInfo: &struct {
			Name    string  `json:"name"`
			Version *string `json:"version,omitempty"`
		}{
			Name:    "test-client",
			Version: stringPtr("1.0.0"),
		},
		Locale:  stringPtr("en-US"),
		RootURI: documentUriPtr("file:///test/workspace"),
		InitializationOptions: map[string]string{
			"key": "value",
		},
		Capabilities: ClientCapabilities_317{},
	}

	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var unmarshaled InitializeParams_316
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if unmarshaled.ProcessID == nil || *unmarshaled.ProcessID != 1234 {
		t.Errorf("Expected ProcessID 1234, got %v", unmarshaled.ProcessID)
	}
	if unmarshaled.ClientInfo.Name != "test-client" {
		t.Errorf("Expected client name 'test-client', got %s", unmarshaled.ClientInfo.Name)
	}
	if unmarshaled.Locale == nil || *unmarshaled.Locale != "en-US" {
		t.Errorf("Expected locale 'en-US', got %v", unmarshaled.Locale)
	}
}

func TestInitializeResult_316(t *testing.T) {
	result := InitializeResult_316{
		Capabilities: ServerCapabilities_316{
			TextDocumentSync: TextDocumentSyncKind(1),
			HoverProvider:    true,
		},
		ServerInfo: &InitializeResultServerInfo{
			Name:    "test-server",
			Version: stringPtr("1.0.0"),
		},
	}

	data, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var unmarshaled InitializeResult_316
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if unmarshaled.ServerInfo.Name != "test-server" {
		t.Errorf("Expected server name 'test-server', got %s", unmarshaled.ServerInfo.Name)
	}
	if unmarshaled.ServerInfo.Version == nil || *unmarshaled.ServerInfo.Version != "1.0.0" {
		t.Errorf("Expected server version '1.0.0', got %v", unmarshaled.ServerInfo.Version)
	}
}

func TestInitializeError(t *testing.T) {
	err := InitializeError{
		Retry: true,
	}

	data, marshalErr := json.Marshal(err)
	if marshalErr != nil {
		t.Fatalf("Marshal failed: %v", marshalErr)
	}

	var unmarshaled InitializeError
	unmarshalErr := json.Unmarshal(data, &unmarshaled)
	if unmarshalErr != nil {
		t.Fatalf("Unmarshal failed: %v", unmarshalErr)
	}

	if !unmarshaled.Retry {
		t.Error("Expected Retry to be true")
	}
}

func TestInitializedParams(t *testing.T) {
	params := InitializedParams{}

	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var unmarshaled InitializedParams
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
}

func TestLogTraceParams(t *testing.T) {
	params := LogTraceParams{
		Message: "test log message",
		Verbose: stringPtr("verbose info"),
	}

	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var unmarshaled LogTraceParams
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if unmarshaled.Message != "test log message" {
		t.Errorf("Expected message 'test log message', got %s", unmarshaled.Message)
	}
	if unmarshaled.Verbose == nil || *unmarshaled.Verbose != "verbose info" {
		t.Errorf("Expected verbose 'verbose info', got %v", unmarshaled.Verbose)
	}
}

func TestSetTraceParams(t *testing.T) {
	params := SetTraceParams{
		Value: TraceValueVerbose,
	}

	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var unmarshaled SetTraceParams
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if unmarshaled.Value != TraceValueVerbose {
		t.Errorf("Expected trace value %v, got %v", TraceValueVerbose, unmarshaled.Value)
	}
}

func TestMethodConstants(t *testing.T) {
	tests := []struct {
		method   string
		expected string
	}{
		{MethodInitialize, "initialize"},
		{MethodInitialized, "initialized"},
		{MethodShutdown, "shutdown"},
		{MethodExit, "exit"},
		{MethodSetTrace, "$/setTrace"},
	}

	for _, test := range tests {
		if test.method != test.expected {
			t.Errorf("Expected method %s to be %s, got %s", test.expected, test.expected, test.method)
		}
	}
}

func TestInitializeErrorCode(t *testing.T) {
	if InitializeErrorCodeUnknownProtocolVersion != 1 {
		t.Errorf("Expected InitializeErrorCodeUnknownProtocolVersion to be 1, got %d", InitializeErrorCodeUnknownProtocolVersion)
	}
}

func stringPtr(s string) *string {
	return &s
}

func documentUriPtr(s string) *DocumentUri {
	uri := DocumentUri(s)
	return &uri
}
