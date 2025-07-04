package protocol

import (
	"sync"

	"github.com/tliron/glsp"
	"github.com/tliron/glsp/protocol/translation"
)

type Handler_317 struct {
	// Base Protocol
	CancelRequest CancelRequestFunc
	Progress      ProgressFunc

	// General Messages
	Initialize  InitializeFunc[InitializeParams_317, InitializeResult_317]
	Initialized InitializedFunc
	Shutdown    ShutdownFunc
	Exit        ExitFunc
	LogTrace    LogTraceFunc
	SetTrace    SetTraceFunc

	// Window
	WindowWorkDoneProgressCancel WindowWorkDoneProgressCancelFunc

	// Workspace
	WorkspaceDidChangeWorkspaceFolders WorkspaceDidChangeWorkspaceFoldersFunc
	WorkspaceDidChangeConfiguration    WorkspaceDidChangeConfigurationFunc
	WorkspaceDidChangeWatchedFiles     WorkspaceDidChangeWatchedFilesFunc
	WorkspaceSymbol                    WorkspaceSymbolFunc
	WorkspaceExecuteCommand            WorkspaceExecuteCommandFunc
	WorkspaceWillCreateFiles           WorkspaceWillCreateFilesFunc
	WorkspaceDidCreateFiles            WorkspaceDidCreateFilesFunc
	WorkspaceWillRenameFiles           WorkspaceWillRenameFilesFunc
	WorkspaceDidRenameFiles            WorkspaceDidRenameFilesFunc
	WorkspaceWillDeleteFiles           WorkspaceWillDeleteFilesFunc
	WorkspaceDidDeleteFiles            WorkspaceDidDeleteFilesFunc
	WorkspaceSemanticTokensRefresh     WorkspaceSemanticTokensRefreshFunc

	// Text Document Synchronization
	TextDocumentDidOpen           TextDocumentDidOpenFunc
	TextDocumentDidChange         TextDocumentDidChangeFunc
	TextDocumentWillSave          TextDocumentWillSaveFunc
	TextDocumentWillSaveWaitUntil TextDocumentWillSaveWaitUntilFunc
	TextDocumentDidSave           TextDocumentDidSaveFunc
	TextDocumentDidClose          TextDocumentDidCloseFunc

	// Language Features
	TextDocumentCompletion              TextDocumentCompletionFunc
	CompletionItemResolve               CompletionItemResolveFunc
	TextDocumentHover                   TextDocumentHoverFunc
	TextDocumentSignatureHelp           TextDocumentSignatureHelpFunc
	TextDocumentDeclaration             TextDocumentDeclarationFunc
	TextDocumentDefinition              TextDocumentDefinitionFunc
	TextDocumentDiagnostic              TextDocumentDiagnosticFunc
	TextDocumentTypeDefinition          TextDocumentTypeDefinitionFunc
	TextDocumentImplementation          TextDocumentImplementationFunc
	TextDocumentReferences              TextDocumentReferencesFunc
	TextDocumentDocumentHighlight       TextDocumentDocumentHighlightFunc
	TextDocumentDocumentSymbol          TextDocumentDocumentSymbolFunc
	TextDocumentCodeAction              TextDocumentCodeActionFunc
	CodeActionResolve                   CodeActionResolveFunc
	TextDocumentCodeLens                TextDocumentCodeLensFunc
	CodeLensResolve                     CodeLensResolveFunc
	TextDocumentDocumentLink            TextDocumentDocumentLinkFunc
	DocumentLinkResolve                 DocumentLinkResolveFunc
	TextDocumentColor                   TextDocumentColorFunc
	TextDocumentColorPresentation       TextDocumentColorPresentationFunc
	TextDocumentFormatting              TextDocumentFormattingFunc
	TextDocumentRangeFormatting         TextDocumentRangeFormattingFunc
	TextDocumentOnTypeFormatting        TextDocumentOnTypeFormattingFunc
	TextDocumentRename                  TextDocumentRenameFunc
	TextDocumentPrepareRename           TextDocumentPrepareRenameFunc
	TextDocumentFoldingRange            TextDocumentFoldingRangeFunc
	TextDocumentSelectionRange          TextDocumentSelectionRangeFunc
	TextDocumentPrepareCallHierarchy    TextDocumentPrepareCallHierarchyFunc
	CallHierarchyIncomingCalls          CallHierarchyIncomingCallsFunc
	CallHierarchyOutgoingCalls          CallHierarchyOutgoingCallsFunc
	TextDocumentSemanticTokensFull      TextDocumentSemanticTokensFullFunc
	TextDocumentSemanticTokensFullDelta TextDocumentSemanticTokensFullDeltaFunc
	TextDocumentSemanticTokensRange     TextDocumentSemanticTokensRangeFunc
	TextDocumentLinkedEditingRange      TextDocumentLinkedEditingRangeFunc
	TextDocumentMoniker                 TextDocumentMonikerFunc

	// Custom Request/Notification
	CustomMethod map[string]CustomMethodHandler

	initialized bool
	lock        sync.Mutex
}

func (h *Handler_317) GetMethodMap() map[Method]glsp.HandlerInterface {

	return map[Method]glsp.HandlerInterface{
		// Base Protocol
		MethodCancelRequest: translation.NewNotificationHandler(h.CancelRequest),
		MethodProgress:      translation.NewNotificationHandler(h.Progress),
		MethodInitialize:    translation.NewTypedHandler(h.Initialize),
		MethodInitialized:   translation.NewNotificationHandler(h.Initialized),
		MethodShutdown:      translation.NewContextOnlyHandler(h.Shutdown),
		MethodExit:          translation.NewContextOnlyHandler(h.Exit),
		MethodLogTrace:      translation.NewNotificationHandler(h.LogTrace),
		MethodSetTrace:      translation.NewNotificationHandler(h.SetTrace),

		// Window
		MethodWindowWorkDoneProgressCancel: translation.NewNotificationHandler(h.WindowWorkDoneProgressCancel),

		// Workspace
		MethodWorkspaceDidChangeWorkspaceFolders: translation.NewNotificationHandler(h.WorkspaceDidChangeWorkspaceFolders),
		MethodWorkspaceDidChangeConfiguration:    translation.NewNotificationHandler(h.WorkspaceDidChangeConfiguration),
		MethodWorkspaceDidChangeWatchedFiles:     translation.NewNotificationHandler(h.WorkspaceDidChangeWatchedFiles),
		MethodWorkspaceSymbol:                    translation.NewTypedHandler(h.WorkspaceSymbol),
		MethodWorkspaceExecuteCommand:            translation.NewTypedHandler(h.WorkspaceExecuteCommand),
		MethodWorkspaceWillCreateFiles:           translation.NewTypedHandler(h.WorkspaceWillCreateFiles),
		MethodWorkspaceDidCreateFiles:            translation.NewNotificationHandler(h.WorkspaceDidCreateFiles),
		MethodWorkspaceWillRenameFiles:           translation.NewTypedHandler(h.WorkspaceWillRenameFiles),
		MethodWorkspaceDidRenameFiles:            translation.NewNotificationHandler(h.WorkspaceDidRenameFiles),
		MethodWorkspaceWillDeleteFiles:           translation.NewTypedHandler(h.WorkspaceWillDeleteFiles),
		MethodWorkspaceDidDeleteFiles:            translation.NewNotificationHandler(h.WorkspaceDidDeleteFiles),
		MethodWorkspaceSemanticTokensRefresh:     translation.NewContextOnlyHandler(h.WorkspaceSemanticTokensRefresh),

		// Text Document Synchronization
		MethodTextDocumentDidOpen:           translation.NewNotificationHandler(h.TextDocumentDidOpen),
		MethodTextDocumentDidChange:         translation.NewNotificationHandler(h.TextDocumentDidChange),
		MethodTextDocumentWillSave:          translation.NewNotificationHandler(h.TextDocumentWillSave),
		MethodTextDocumentWillSaveWaitUntil: translation.NewTypedHandler(h.TextDocumentWillSaveWaitUntil),
		MethodTextDocumentDidSave:           translation.NewNotificationHandler(h.TextDocumentDidSave),
		MethodTextDocumentDidClose:          translation.NewNotificationHandler(h.TextDocumentDidClose),

		// Language Features
		MethodTextDocumentCompletion:              translation.NewTypedHandler(h.TextDocumentCompletion),
		MethodCompletionItemResolve:               translation.NewTypedHandler(h.CompletionItemResolve),
		MethodTextDocumentHover:                   translation.NewTypedHandler(h.TextDocumentHover),
		MethodTextDocumentSignatureHelp:           translation.NewTypedHandler(h.TextDocumentSignatureHelp),
		MethodTextDocumentDeclaration:             translation.NewTypedHandler(h.TextDocumentDeclaration),
		MethodTextDocumentDefinition:              translation.NewTypedHandler(h.TextDocumentDefinition),
		MethodTextDocumentTypeDefinition:          translation.NewTypedHandler(h.TextDocumentTypeDefinition),
		MethodTextDocumentImplementation:          translation.NewTypedHandler(h.TextDocumentImplementation),
		MethodTextDocumentReferences:              translation.NewTypedHandler(h.TextDocumentReferences),
		MethodTextDocumentDocumentHighlight:       translation.NewTypedHandler(h.TextDocumentDocumentHighlight),
		MethodTextDocumentDocumentSymbol:          translation.NewTypedHandler(h.TextDocumentDocumentSymbol),
		MethodTextDocumentCodeAction:              translation.NewTypedHandler(h.TextDocumentCodeAction),
		MethodCodeActionResolve:                   translation.NewTypedHandler(h.CodeActionResolve),
		MethodTextDocumentCodeLens:                translation.NewTypedHandler(h.TextDocumentCodeLens),
		MethodCodeLensResolve:                     translation.NewTypedHandler(h.CodeLensResolve),
		MethodTextDocumentDocumentLink:            translation.NewTypedHandler(h.TextDocumentDocumentLink),
		MethodDocumentLinkResolve:                 translation.NewTypedHandler(h.DocumentLinkResolve),
		MethodTextDocumentColor:                   translation.NewTypedHandler(h.TextDocumentColor),
		MethodTextDocumentColorPresentation:       translation.NewTypedHandler(h.TextDocumentColorPresentation),
		MethodTextDocumentFormatting:              translation.NewTypedHandler(h.TextDocumentFormatting),
		MethodTextDocumentRangeFormatting:         translation.NewTypedHandler(h.TextDocumentRangeFormatting),
		MethodTextDocumentOnTypeFormatting:        translation.NewTypedHandler(h.TextDocumentOnTypeFormatting),
		MethodTextDocumentRename:                  translation.NewTypedHandler(h.TextDocumentRename),
		MethodTextDocumentPrepareRename:           translation.NewTypedHandler(h.TextDocumentPrepareRename),
		MethodTextDocumentFoldingRange:            translation.NewTypedHandler(h.TextDocumentFoldingRange),
		MethodTextDocumentSelectionRange:          translation.NewTypedHandler(h.TextDocumentSelectionRange),
		MethodTextDocumentPrepareCallHierarchy:    translation.NewTypedHandler(h.TextDocumentPrepareCallHierarchy),
		MethodCallHierarchyIncomingCalls:          translation.NewTypedHandler(h.CallHierarchyIncomingCalls),
		MethodCallHierarchyOutgoingCalls:          translation.NewTypedHandler(h.CallHierarchyOutgoingCalls),
		MethodTextDocumentSemanticTokensFull:      translation.NewTypedHandler(h.TextDocumentSemanticTokensFull),
		MethodTextDocumentSemanticTokensFullDelta: translation.NewTypedHandler(h.TextDocumentSemanticTokensFullDelta),
		MethodTextDocumentSemanticTokensRange:     translation.NewTypedHandler(h.TextDocumentSemanticTokensRange),
		MethodTextDocumentLinkedEditingRange:      translation.NewTypedHandler(h.TextDocumentLinkedEditingRange),
		MethodTextDocumentMoniker:                 translation.NewTypedHandler(h.TextDocumentMoniker),
		MethodTextDocumentDiagnostic:              translation.NewTypedHandler(h.TextDocumentDiagnostic),
	}
}

func (h *Handler_317) GetCustomMethods() map[string]glsp.HandlerInterface {
	var customMethods = make(map[string]glsp.HandlerInterface)
	for method, handler := range h.CustomMethod {
		customMethods[method] = translation.NewCustomHandler(handler.Func)
	}
	return customMethods
}

func (h *Handler_317) IsInitialized() bool {
	h.lock.Lock()
	defer h.lock.Unlock()
	return h.initialized
}

func (h *Handler_317) SetInitialized(initialized bool) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.initialized = initialized
}

func (h *Handler_317) CreateServerCapabilities() ServerCapabilities_317 {
	var capabilities ServerCapabilities_317

	if (h.TextDocumentDidOpen != nil) || (h.TextDocumentDidClose != nil) {
		if _, ok := capabilities.TextDocumentSync.(*TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).OpenClose = &True
	}

	if h.TextDocumentDidChange != nil {
		if _, ok := capabilities.TextDocumentSync.(*TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &TextDocumentSyncOptions{}
		}
		// This can be overriden to TextDocumentSyncKindFull
		value := TextDocumentSyncKindIncremental
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).Change = &value
	}

	if h.TextDocumentWillSave != nil {
		if _, ok := capabilities.TextDocumentSync.(*TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).WillSave = &True
	}

	if h.TextDocumentWillSaveWaitUntil != nil {
		if _, ok := capabilities.TextDocumentSync.(*TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).WillSaveWaitUntil = &True
	}

	if h.TextDocumentDidSave != nil {
		if _, ok := capabilities.TextDocumentSync.(*TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).Save = &True
	}

	if h.TextDocumentCompletion != nil {
		capabilities.CompletionProvider = &CompletionOptions{}
	}

	if h.TextDocumentHover != nil {
		capabilities.HoverProvider = true
	}

	if h.TextDocumentSignatureHelp != nil {
		capabilities.SignatureHelpProvider = &SignatureHelpOptions{}
	}

	if h.TextDocumentDeclaration != nil {
		capabilities.DeclarationProvider = true
	}

	if h.TextDocumentDefinition != nil {
		capabilities.DefinitionProvider = true
	}

	if h.TextDocumentTypeDefinition != nil {
		capabilities.TypeDefinitionProvider = true
	}

	if h.TextDocumentImplementation != nil {
		capabilities.ImplementationProvider = true
	}

	if h.TextDocumentReferences != nil {
		capabilities.ReferencesProvider = true
	}

	if h.TextDocumentDocumentHighlight != nil {
		capabilities.DocumentHighlightProvider = true
	}

	if h.TextDocumentDocumentSymbol != nil {
		capabilities.DocumentSymbolProvider = true
	}

	if h.TextDocumentCodeAction != nil {
		capabilities.CodeActionProvider = true
	}

	if h.TextDocumentCodeLens != nil {
		capabilities.CodeLensProvider = &CodeLensOptions{}
	}

	if h.TextDocumentDocumentLink != nil {
		capabilities.DocumentLinkProvider = &DocumentLinkOptions{}
	}

	if h.TextDocumentColor != nil {
		capabilities.ColorProvider = true
	}

	if h.TextDocumentFormatting != nil {
		capabilities.DocumentFormattingProvider = true
	}

	if h.TextDocumentRangeFormatting != nil {
		capabilities.DocumentRangeFormattingProvider = true
	}

	if h.TextDocumentOnTypeFormatting != nil {
		capabilities.DocumentOnTypeFormattingProvider = &DocumentOnTypeFormattingOptions{}
	}

	if h.TextDocumentRename != nil {
		capabilities.RenameProvider = true
	}

	if h.TextDocumentFoldingRange != nil {
		capabilities.FoldingRangeProvider = true
	}

	if h.WorkspaceExecuteCommand != nil {
		capabilities.ExecuteCommandProvider = &ExecuteCommandOptions{}
	}

	if h.TextDocumentSelectionRange != nil {
		capabilities.SelectionRangeProvider = true
	}

	if h.TextDocumentLinkedEditingRange != nil {
		capabilities.LinkedEditingRangeProvider = true
	}

	if h.TextDocumentPrepareCallHierarchy != nil {
		capabilities.CallHierarchyProvider = true
	}

	if h.TextDocumentSemanticTokensFull != nil {
		if _, ok := capabilities.SemanticTokensProvider.(*SemanticTokensOptions); !ok {
			capabilities.SemanticTokensProvider = &SemanticTokensOptions{}
		}
		if h.TextDocumentSemanticTokensFullDelta != nil {
			capabilities.SemanticTokensProvider.(*SemanticTokensOptions).Full = &SemanticDelta{}
			capabilities.SemanticTokensProvider.(*SemanticTokensOptions).Full.(*SemanticDelta).Delta = &True
		} else {
			capabilities.SemanticTokensProvider.(*SemanticTokensOptions).Full = true
		}
	}

	if h.TextDocumentSemanticTokensRange != nil {
		if _, ok := capabilities.SemanticTokensProvider.(*SemanticTokensOptions); !ok {
			capabilities.SemanticTokensProvider = &SemanticTokensOptions{}
		}
		capabilities.SemanticTokensProvider.(*SemanticTokensOptions).Range = true
	}

	// TODO: h.TextDocumentSemanticTokensRefresh?

	if h.TextDocumentMoniker != nil {
		capabilities.MonikerProvider = true
	}

	if h.WorkspaceSymbol != nil {
		capabilities.WorkspaceSymbolProvider = true
	}

	if h.WorkspaceDidCreateFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidCreate = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if h.WorkspaceWillCreateFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillCreate = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if h.WorkspaceDidRenameFiles != nil {
		capabilities.RenameProvider = true
		if capabilities.Workspace == nil {
			capabilities.Workspace = &ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidRename = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if h.WorkspaceWillRenameFiles != nil {
		capabilities.RenameProvider = true
		if capabilities.Workspace == nil {
			capabilities.Workspace = &ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillRename = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if h.WorkspaceDidDeleteFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidDelete = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if h.WorkspaceWillDeleteFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillDelete = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if h.TextDocumentDiagnostic != nil {
		capabilities.DiagnosticProvider = DiagnosticOptions{
			InterFileDependencies: true,
			WorkspaceDiagnostics:  false,
		}
	}

	return capabilities
}

func (h *Handler_317) GetVersion() glsp.LspProtocolVersion {
	return glsp.Protocol_3_17
}
