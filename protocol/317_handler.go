//go:generate go run ../tools/genmethodmap/main.go 317_handler

package protocol

type Handler_317 struct {
	// Base Protocol
	CancelRequest NotificationFunc[CancelParams]
	Progress      NotificationFunc[ProgressParams]

	// General Messages
	Exit        ContextOnlyFunc
	Initialize  RequestFunc[InitializeParams_316, InitializeResult_316]
	Initialized NotificationFunc[InitializedParams]
	SetTrace    NotificationFunc[SetTraceParams]
	Shutdown    ContextOnlyFunc

	// Window
	WindowWorkDoneProgressCancel NotificationFunc[WorkDoneProgressCancelParams]

	// Workspace
	WorkspaceDidChangeConfiguration    NotificationFunc[DidChangeConfigurationParams]
	WorkspaceDidChangeWatchedFiles     NotificationFunc[DidChangeWatchedFilesParams]
	WorkspaceDidChangeWorkspaceFolders NotificationFunc[DidChangeWorkspaceFoldersParams]
	WorkspaceDidCreateFiles            NotificationFunc[CreateFilesParams]
	WorkspaceDidDeleteFiles            NotificationFunc[DeleteFilesParams]
	WorkspaceDidRenameFiles            NotificationFunc[RenameFilesParams]
	WorkspaceExecuteCommand            RequestFunc[ExecuteCommandParams, any]
	WorkspaceSymbol                    RequestFunc[WorkspaceSymbolParams, []SymbolInformation]
	WorkspaceWillCreateFiles           RequestFunc[CreateFilesParams, WorkspaceEdit]
	WorkspaceWillDeleteFiles           RequestFunc[DeleteFilesParams, WorkspaceEdit]
	WorkspaceWillRenameFiles           RequestFunc[RenameFilesParams, WorkspaceEdit]

	// Text Document Synchronization
	TextDocumentDidChange         NotificationFunc[DidChangeTextDocumentParams]
	TextDocumentDidClose          NotificationFunc[DidCloseTextDocumentParams]
	TextDocumentDidOpen           NotificationFunc[DidOpenTextDocumentParams]
	TextDocumentDidSave           NotificationFunc[DidSaveTextDocumentParams]
	TextDocumentWillSave          NotificationFunc[WillSaveTextDocumentParams]
	TextDocumentWillSaveWaitUntil RequestFunc[WillSaveTextDocumentParams, []TextEdit]

	// Language Features
	CallHierarchyIncomingCalls          RequestFunc[CallHierarchyIncomingCallsParams, []CallHierarchyIncomingCall]
	CallHierarchyOutgoingCalls          RequestFunc[CallHierarchyOutgoingCallsParams, []CallHierarchyOutgoingCall]
	CodeActionResolve                   RequestFunc[CodeAction, CodeAction]
	CodeLensResolve                     RequestFunc[CodeLens, CodeLens]
	CompletionItemResolve               RequestFunc[CompletionItem, CompletionItem]
	DocumentLinkResolve                 RequestFunc[DocumentLink, DocumentLink]
	TextDocumentCodeAction              RequestFunc[CodeAction, any]
	TextDocumentCodeLens                RequestFunc[CodeLensParams, []CodeLens]
	TextDocumentDocumentColor           RequestFunc[DocumentColorParams, []ColorInformation]
	TextDocumentColorPresentation       RequestFunc[ColorPresentationParams, []ColorPresentation]
	TextDocumentCompletion              RequestFunc[CompletionParams, any]
	TextDocumentDeclaration             RequestFunc[DeclarationParams, any]
	TextDocumentDefinition              RequestFunc[DefinitionParams, any]
	TextDocumentDiagnostic              RequestFunc[DocumentDiagnosticParams, any]
	TextDocumentDocumentHighlight       RequestFunc[DocumentHighlightParams, []DocumentHighlight]
	TextDocumentDocumentLink            RequestFunc[DocumentLinkParams, []DocumentLink]
	TextDocumentDocumentSymbol          RequestFunc[DocumentSymbolParams, any]
	TextDocumentFoldingRange            RequestFunc[FoldingRangeParams, []FoldingRange]
	TextDocumentFormatting              RequestFunc[DocumentFormattingParams, []TextEdit]
	TextDocumentHover                   RequestFunc[HoverParams, Hover]
	TextDocumentImplementation          RequestFunc[ImplementationParams, any]
	TextDocumentLinkedEditingRange      RequestFunc[LinkedEditingRangeParams, LinkedEditingRanges]
	TextDocumentMoniker                 RequestFunc[MonikerParams, []Moniker]
	TextDocumentOnTypeFormatting        RequestFunc[DocumentOnTypeFormattingParams, []TextEdit]
	TextDocumentPrepareCallHierarchy    RequestFunc[CallHierarchyPrepareParams, []CallHierarchyItem]
	TextDocumentPrepareRename           RequestFunc[PrepareRenameParams, any]
	TextDocumentRangeFormatting         RequestFunc[DocumentRangeFormattingParams, []TextEdit]
	TextDocumentReferences              RequestFunc[ReferenceParams, []Location]
	TextDocumentRename                  RequestFunc[RenameParams, WorkspaceEdit]
	TextDocumentSelectionRange          RequestFunc[SelectionRangeParams, []SelectionRange]
	TextDocumentSemanticTokensFull      RequestFunc[SemanticTokensParams, SemanticTokens]
	TextDocumentSemanticTokensFullDelta RequestFunc[SemanticTokensDeltaParams, any]
	TextDocumentSemanticTokensRange     RequestFunc[SemanticTokensRangeParams, SemanticTokens]
	TextDocumentSignatureHelp           RequestFunc[SignatureHelpParams, SignatureHelp]
	TextDocumentTypeDefinition          RequestFunc[TypeDefinitionParams, any]

	common_handler
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

	if h.TextDocumentDocumentColor != nil {
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
