package protocol

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestDocumentChangeValid(t *testing.T) {
	if !(&DocumentChange{CreateFile: &CreateFile{}}).Valid() {
		t.Fatal("Expected one populated variant to be valid")
	}
	if (&DocumentChange{}).Valid() {
		t.Fatal("Expected empty DocumentChange to be invalid")
	}
	if (&DocumentChange{CreateFile: &CreateFile{}, DeleteFile: &DeleteFile{}}).Valid() {
		t.Fatal("Expected multiple populated variants to be invalid")
	}
}

func TestDocumentChangeUnmarshalTextDocumentEdit(t *testing.T) {
	var change DocumentChange
	err := json.Unmarshal([]byte(`{"textDocument":{"uri":"file:///tmp/a.go","version":1},"edits":[]}`), &change)
	if err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}

	if change.TextDocumentEdit == nil {
		t.Fatal("Expected text document edit variant")
	}
	if change.TextDocumentEdit.TextDocument.URI != "file:///tmp/a.go" {
		t.Fatalf("Expected text document URI to unmarshal, got %q", change.TextDocumentEdit.TextDocument.URI)
	}
}

func TestDocumentChangeUnmarshalFileOperations(t *testing.T) {
	tests := []struct {
		name string
		json string
		want func(DocumentChange) bool
	}{
		{
			name: "create",
			json: `{"kind":"create","uri":"file:///tmp/a.go"}`,
			want: func(change DocumentChange) bool {
				return change.CreateFile != nil && change.CreateFile.URI == "file:///tmp/a.go"
			},
		},
		{
			name: "rename",
			json: `{"kind":"rename","oldUri":"file:///tmp/a.go","newUri":"file:///tmp/b.go"}`,
			want: func(change DocumentChange) bool {
				return change.RenameFile != nil && change.RenameFile.NewURI == "file:///tmp/b.go"
			},
		},
		{
			name: "delete",
			json: `{"kind":"delete","uri":"file:///tmp/a.go"}`,
			want: func(change DocumentChange) bool {
				return change.DeleteFile != nil && change.DeleteFile.URI == "file:///tmp/a.go"
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var change DocumentChange
			if err := json.Unmarshal([]byte(test.json), &change); err != nil {
				t.Fatalf("Unmarshal returned error: %v", err)
			}
			if !test.want(change) {
				t.Fatalf("Unexpected variant after unmarshal: %#v", change)
			}
		})
	}
}

func TestDocumentChangeUnmarshalRejectsUnknownKind(t *testing.T) {
	var change DocumentChange
	err := json.Unmarshal([]byte(`{"kind":"copy","uri":"file:///tmp/a.go"}`), &change)
	if err == nil {
		t.Fatal("Expected unknown kind to fail")
	}
	if !strings.Contains(err.Error(), "unexpected kind") {
		t.Fatalf("Expected unexpected kind error, got %v", err)
	}
}

func TestDocumentChangeMarshalFileOperation(t *testing.T) {
	change := &DocumentChange{CreateFile: &CreateFile{Kind: "create", URI: "file:///tmp/a.go"}}

	data, err := json.Marshal(change)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}

	if string(data) != `{"kind":"create","uri":"file:///tmp/a.go"}` {
		t.Fatalf("Unexpected marshal output: %s", data)
	}
}

func TestDocumentChangeMarshalRejectsEmptyUnion(t *testing.T) {
	_, err := json.Marshal(&DocumentChange{})
	if err == nil {
		t.Fatal("Expected empty DocumentChange to fail marshal")
	}
}
