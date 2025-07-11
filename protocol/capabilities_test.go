package protocol

import (
	"github.com/tliron/glsp"
	"testing"
)

func TestCreateServerCapabilities_TextDocumentDidOpen(t *testing.T) {
	handler := &Handler_316{
		TextDocumentDidOpen: NotificationFunc[DidOpenTextDocumentParams](func(_ *glsp.Context, _ *DidOpenTextDocumentParams) error { return nil }),
	}
	caps := CreateServerCapabilities(handler)

	sync, ok := caps.TextDocumentSync.(*TextDocumentSyncOptions)
	if !ok {
		t.Fatalf("TextDocumentSync should be *TextDocumentSyncOptions, got %T", caps.TextDocumentSync)
	}
	if sync.OpenClose != true {
		t.Errorf("OpenClose should be true when TextDocumentDidOpen is set")
	}

	if caps.HoverProvider != nil {
		t.Errorf("HoverProvider should be nil when TextDocumentHover is not set")
	}
}

func TestCreateServerCapabilities_MultipleHandlers(t *testing.T) {
	handler := &Handler_316{
		TextDocumentDidOpen: NotificationFunc[DidOpenTextDocumentParams](func(_ *glsp.Context, _ *DidOpenTextDocumentParams) error { return nil }),
		TextDocumentHover:   RequestFunc[HoverParams, Hover](func(_ *glsp.Context, _ *HoverParams) (Hover, error) { return Hover{}, nil }),
	}
	caps := CreateServerCapabilities(handler)

	sync, ok := caps.TextDocumentSync.(*TextDocumentSyncOptions)
	if !ok {
		t.Fatalf("TextDocumentSync should be *TextDocumentSyncOptions, got %T", caps.TextDocumentSync)
	}
	if sync.OpenClose != true {
		t.Errorf("OpenClose should be true when TextDocumentDidOpen is set")
	}

	if caps.HoverProvider == nil {
		t.Errorf("HoverProvider should not be nil when TextDocumentHover is set")
	}
}

func TestCreateServerCapabilities_NoHandlers(t *testing.T) {
	handler := &Handler_316{}
	caps := CreateServerCapabilities(handler)

	if caps.TextDocumentSync != nil {
		t.Errorf("TextDocumentSync should be nil when no handlers are set")
	}
	if caps.HoverProvider != nil {
		t.Errorf("HoverProvider should be nil when no handlers are set")
	}
}

func TestCreateServerCapabilities_PointerVsValue(t *testing.T) {
	// Handler as value
	handlerVal := Handler_316{
		TextDocumentDidOpen: NotificationFunc[DidOpenTextDocumentParams](func(_ *glsp.Context, _ *DidOpenTextDocumentParams) error { return nil }),
	}
	capsVal := CreateServerCapabilities(&handlerVal)
	syncVal, ok := capsVal.TextDocumentSync.(*TextDocumentSyncOptions)
	if !ok {
		t.Fatalf("TextDocumentSync (value) should be *TextDocumentSyncOptions, got %T", capsVal.TextDocumentSync)
	}
	if syncVal.OpenClose != true {
		t.Errorf("OpenClose (value) should be true when TextDocumentDidOpen is set")
	}

	// Handler as pointer
	handlerPtr := &Handler_316{
		TextDocumentDidOpen: NotificationFunc[DidOpenTextDocumentParams](func(_ *glsp.Context, _ *DidOpenTextDocumentParams) error { return nil }),
	}
	capsPtr := CreateServerCapabilities(handlerPtr)
	syncPtr, ok := capsPtr.TextDocumentSync.(*TextDocumentSyncOptions)
	if !ok {
		t.Fatalf("TextDocumentSync (pointer) should be *TextDocumentSyncOptions, got %T", capsPtr.TextDocumentSync)
	}
	if syncPtr.OpenClose != true {
		t.Errorf("OpenClose (pointer) should be true when TextDocumentDidOpen is set")
	}
}
