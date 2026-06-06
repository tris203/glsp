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

func TestCreateServerCapabilities_PointerBackedProviders(t *testing.T) {
	handler := &Handler_316{
		TextDocumentCompletion:   RequestFunc[CompletionParams, CompletionList](func(_ *glsp.Context, _ *CompletionParams) (CompletionList, error) { return CompletionList{}, nil }),
		TextDocumentCodeLens:     RequestFunc[CodeLensParams, []CodeLens](func(_ *glsp.Context, _ *CodeLensParams) ([]CodeLens, error) { return nil, nil }),
		TextDocumentDocumentLink: RequestFunc[DocumentLinkParams, []DocumentLink](func(_ *glsp.Context, _ *DocumentLinkParams) ([]DocumentLink, error) { return nil, nil }),
		TextDocumentDidSave:      NotificationFunc[DidSaveTextDocumentParams](func(_ *glsp.Context, _ *DidSaveTextDocumentParams) error { return nil }),
	}

	caps := CreateServerCapabilities(handler)
	if caps.CompletionProvider == nil {
		t.Error("CompletionProvider should be set")
	}
	if caps.CodeLensProvider == nil {
		t.Error("CodeLensProvider should be set")
	}
	if caps.DocumentLinkProvider == nil {
		t.Error("DocumentLinkProvider should be set")
	}
	sync, ok := caps.TextDocumentSync.(*TextDocumentSyncOptions)
	if !ok {
		t.Fatalf("TextDocumentSync should be *TextDocumentSyncOptions, got %T", caps.TextDocumentSync)
	}
	if sync.Save == nil {
		t.Error("Save should be set")
	}
}

func TestCreateServerCapabilities_DiagnosticProvider(t *testing.T) {
	handler := &Handler_317{
		TextDocumentDiagnostic: RequestFunc[DocumentDiagnosticParams, DocumentDiagnosticReport](func(_ *glsp.Context, _ *DocumentDiagnosticParams) (DocumentDiagnosticReport, error) {
			return DocumentDiagnosticReport{}, nil
		}),
	}

	caps := CreateServerCapabilities(handler)
	if caps.DiagnosticProvider == nil {
		t.Fatal("DiagnosticProvider should be set")
	}
	options, ok := caps.DiagnosticProvider.Value.(DiagnosticOptions)
	if !ok {
		t.Fatalf("DiagnosticProvider should wrap DiagnosticOptions, got %T", caps.DiagnosticProvider.Value)
	}
	if !options.InterFileDependencies {
		t.Error("DiagnosticProvider should set InterFileDependencies")
	}
}

func TestCreateServerCapabilities_SemanticTokensDelta(t *testing.T) {
	handler := &Handler_316{
		TextDocumentSemanticTokensFull:      RequestFunc[SemanticTokensParams, SemanticTokens](func(_ *glsp.Context, _ *SemanticTokensParams) (SemanticTokens, error) { return SemanticTokens{}, nil }),
		TextDocumentSemanticTokensFullDelta: RequestFunc[SemanticTokensDeltaParams, any](func(_ *glsp.Context, _ *SemanticTokensDeltaParams) (any, error) { return nil, nil }),
		TextDocumentSemanticTokensRange: RequestFunc[SemanticTokensRangeParams, SemanticTokens](func(_ *glsp.Context, _ *SemanticTokensRangeParams) (SemanticTokens, error) {
			return SemanticTokens{}, nil
		}),
	}

	caps := CreateServerCapabilities(handler)
	options, ok := caps.SemanticTokensProvider.(*SemanticTokensOptions)
	if !ok {
		t.Fatalf("SemanticTokensProvider should be *SemanticTokensOptions, got %T", caps.SemanticTokensProvider)
	}
	if options.Full == nil {
		t.Fatal("SemanticTokensProvider.Full should be set")
	}
	delta, ok := options.Full.Value.(SemanticTokensFullDelta)
	if !ok {
		t.Fatalf("SemanticTokensProvider.Full should wrap SemanticTokensFullDelta, got %T", options.Full.Value)
	}
	if !delta.Delta {
		t.Error("SemanticTokensProvider.Full delta should be true")
	}
	if options.Range == nil || options.Range.Value != true {
		t.Errorf("SemanticTokensProvider.Range should be true, got %#v", options.Range)
	}
}
