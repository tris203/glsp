package protocol

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestParseDocumentURI_CanonicalizesFileURI(t *testing.T) {
	uri, err := ParseDocumentURI("file:///tmp/a%20b.go")
	if err != nil {
		t.Fatalf("ParseDocumentURI returned error: %v", err)
	}

	if uri != "file:///tmp/a%20b.go" {
		t.Fatalf("Expected canonical escaped URI, got %q", uri)
	}
	if uri.Path() != filepath.FromSlash("/tmp/a b.go") {
		t.Fatalf("Expected unescaped file path, got %q", uri.Path())
	}
}

func TestParseDocumentURI_AcceptsVSCodeTwoSlashFileURI(t *testing.T) {
	uri, err := ParseDocumentURI("file://tmp/a.go")
	if err != nil {
		t.Fatalf("ParseDocumentURI returned error: %v", err)
	}

	if uri != "file:///tmp/a.go" {
		t.Fatalf("Expected two-slash URI to be normalized, got %q", uri)
	}
}

func TestParseDocumentURI_NormalizesWindowsDriveLetter(t *testing.T) {
	uri, err := ParseDocumentURI("file:///c:/project/readme.md")
	if err != nil {
		t.Fatalf("ParseDocumentURI returned error: %v", err)
	}

	if uri != "file:///C:/project/readme.md" {
		t.Fatalf("Expected uppercase Windows drive letter, got %q", uri)
	}
}

func TestParseDocumentURI_RejectsNonFileScheme(t *testing.T) {
	_, err := ParseDocumentURI("https://example.com/file.go")
	if err == nil {
		t.Fatal("Expected non-file URI to be rejected")
	}
}

func TestDocumentURIPathHelpers(t *testing.T) {
	uri := DocumentURI("file:///tmp/project/main.go")

	if uri.Base() != "main.go" {
		t.Fatalf("Expected base main.go, got %q", uri.Base())
	}
	if uri.DirPath() != filepath.FromSlash("/tmp/project") {
		t.Fatalf("Expected dir path /tmp/project, got %q", uri.DirPath())
	}
	if uri.Dir() != "file:///tmp/project" {
		t.Fatalf("Expected dir URI, got %q", uri.Dir())
	}
}

func TestURIFromPath(t *testing.T) {
	uri := URIFromPath(filepath.FromSlash("/tmp/a b.go"))

	if uri != "file:///tmp/a%20b.go" {
		t.Fatalf("Expected escaped absolute file URI, got %q", uri)
	}
	if URIFromPath("") != "" {
		t.Fatal("Expected empty path to produce empty URI")
	}
}

func TestDocumentURIUnmarshalText(t *testing.T) {
	var uri DocumentURI
	if err := uri.UnmarshalText([]byte("file://tmp/a.go")); err != nil {
		t.Fatalf("UnmarshalText returned error: %v", err)
	}

	if !strings.HasPrefix(string(uri), "file:///") {
		t.Fatalf("Expected normalized file URI, got %q", uri)
	}
}
