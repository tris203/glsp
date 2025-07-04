package protocol

import (
	"encoding/json"
	"testing"
)

func TestIntegerOrString_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		expected string
	}{
		{"integer", Integer(42), "42"},
		{"string", "test", `"test"`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ios := IntegerOrString{Value: test.value}
			result, err := ios.MarshalJSON()
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}
			if string(result) != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, string(result))
			}
		})
	}
}

func TestIntegerOrString_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{"integer", "42", Integer(42)},
		{"string", `"test"`, "test"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var ios IntegerOrString
			err := json.Unmarshal([]byte(test.input), &ios)
			if err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}
			if ios.Value != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, ios.Value)
			}
		})
	}
}

func TestIntegerOrString_UnmarshalJSON_Invalid(t *testing.T) {
	var ios IntegerOrString
	err := json.Unmarshal([]byte("true"), &ios)
	if err == nil {
		t.Error("Expected error for invalid input, got nil")
	}
}

func TestBoolOrString_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		expected string
	}{
		{"bool_true", true, "true"},
		{"bool_false", false, "false"},
		{"string", "test", `"test"`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bos := BoolOrString{Value: test.value}
			result, err := bos.MarshalJSON()
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}
			if string(result) != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, string(result))
			}
		})
	}
}

func TestBoolOrString_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{"bool_true", "true", true},
		{"bool_false", "false", false},
		{"string", `"test"`, "test"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var bos BoolOrString
			err := json.Unmarshal([]byte(test.input), &bos)
			if err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}
			if bos.Value != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, bos.Value)
			}
		})
	}
}

func TestBoolOrString_UnmarshalJSON_Invalid(t *testing.T) {
	var bos BoolOrString
	err := json.Unmarshal([]byte("42"), &bos)
	if err == nil {
		t.Error("Expected error for invalid input, got nil")
	}
}

func TestBoolOrString_String(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		expected string
	}{
		{"bool_true", true, "true"},
		{"bool_false", false, "false"},
		{"string", "test", "test"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bos := BoolOrString{Value: test.value}
			result := bos.String()
			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}

func TestCancelParams(t *testing.T) {
	// Test that CancelParams can be created and the ID field is accessible
	params := CancelParams{
		ID: IntegerOrString{Value: Integer(123)},
	}

	if params.ID.Value != Integer(123) {
		t.Errorf("Expected ID 123, got %v", params.ID.Value)
	}

	// Test with string ID
	params2 := CancelParams{
		ID: IntegerOrString{Value: "test-id"},
	}

	if params2.ID.Value != "test-id" {
		t.Errorf("Expected ID 'test-id', got %v", params2.ID.Value)
	}
}
func TestProgressParams(t *testing.T) {
	// Test that ProgressParams can be created and the Token field is accessible
	params := ProgressParams{
		Token: ProgressToken{Value: "test-token"},
		Value: map[string]string{"message": "progress update"},
	}

	if params.Token.Value != "test-token" {
		t.Errorf("Expected token 'test-token', got %v", params.Token.Value)
	}

	if params.Value == nil {
		t.Error("Expected value to be set")
	}
}
func TestConstants(t *testing.T) {
	if MethodCancelRequest != "$/cancelRequest" {
		t.Errorf("Expected MethodCancelRequest to be '$/cancelRequest', got %s", MethodCancelRequest)
	}
	if MethodProgress != "$/progress" {
		t.Errorf("Expected MethodProgress to be '$/progress', got %s", MethodProgress)
	}
}

func TestBoolConstants(t *testing.T) {
	if True != true {
		t.Error("True constant should be true")
	}
	if False != false {
		t.Error("False constant should be false")
	}
}
