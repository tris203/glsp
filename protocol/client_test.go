package protocol

import (
	"encoding/json"
	"testing"
)

func TestRegistration(t *testing.T) {
	reg := Registration{
		ID:              "test-id",
		Method:          "textDocument/hover",
		RegisterOptions: map[string]any{"documentSelector": []string{"*.go"}},
	}

	data, err := json.Marshal(reg)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var unmarshaled Registration
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if unmarshaled.ID != "test-id" {
		t.Errorf("Expected ID 'test-id', got %s", unmarshaled.ID)
	}
	if unmarshaled.Method != "textDocument/hover" {
		t.Errorf("Expected method 'textDocument/hover', got %s", unmarshaled.Method)
	}
}

func TestRegistrationParams(t *testing.T) {
	params := RegistrationParams{
		Registrations: []Registration{
			{
				ID:     "reg1",
				Method: "textDocument/hover",
			},
			{
				ID:     "reg2",
				Method: "textDocument/completion",
			},
		},
	}

	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var unmarshaled RegistrationParams
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if len(unmarshaled.Registrations) != 2 {
		t.Errorf("Expected 2 registrations, got %d", len(unmarshaled.Registrations))
	}
	if unmarshaled.Registrations[0].ID != "reg1" {
		t.Errorf("Expected first registration ID 'reg1', got %s", unmarshaled.Registrations[0].ID)
	}
}

func TestUnregistration(t *testing.T) {
	unreg := Unregistration{
		ID:     "test-id",
		Method: "textDocument/hover",
	}

	data, err := json.Marshal(unreg)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var unmarshaled Unregistration
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if unmarshaled.ID != "test-id" {
		t.Errorf("Expected ID 'test-id', got %s", unmarshaled.ID)
	}
	if unmarshaled.Method != "textDocument/hover" {
		t.Errorf("Expected method 'textDocument/hover', got %s", unmarshaled.Method)
	}
}

func TestUnregistrationParams(t *testing.T) {
	params := UnregistrationParams{
		Unregisterations: []Unregistration{
			{
				ID:     "unreg1",
				Method: "textDocument/hover",
			},
		},
	}

	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var unmarshaled UnregistrationParams
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if len(unmarshaled.Unregisterations) != 1 {
		t.Errorf("Expected 1 unregistration, got %d", len(unmarshaled.Unregisterations))
	}
	if unmarshaled.Unregisterations[0].ID != "unreg1" {
		t.Errorf("Expected unregistration ID 'unreg1', got %s", unmarshaled.Unregisterations[0].ID)
	}
}

func TestClientConstants(t *testing.T) {
	if ServerClientRegisterCapability != "client/registerCapability" {
		t.Errorf("Expected ServerClientRegisterCapability to be 'client/registerCapability', got %s", ServerClientRegisterCapability)
	}
	if ServerClientUnregisterCapability != "client/unregisterCapability" {
		t.Errorf("Expected ServerClientUnregisterCapability to be 'client/unregisterCapability', got %s", ServerClientUnregisterCapability)
	}
}
