//go:generate go run ../tools/genmethodmap/main.go 317_handler

package protocol

type Handler_317 struct {
	// Base Protocol
	CancelRequest NotificationFunc[CancelParams]
	Progress      NotificationFunc[ProgressParams]

	// General Messages
	Exit        ContextOnlyFunc
	Initialize  RequestFunc[InitializeParams, InitializeResult]
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
	TextDocumentCodeAction              RequestFunc[CodeActionParams, []CodeAction]
	TextDocumentCodeLens                RequestFunc[CodeLensParams, []CodeLens]
	TextDocumentDocumentColor           RequestFunc[DocumentColorParams, []ColorInformation]
	TextDocumentColorPresentation       RequestFunc[ColorPresentationParams, []ColorPresentation]
	TextDocumentCompletion              RequestFunc[CompletionParams, CompletionList]
	TextDocumentDeclaration             RequestFunc[DeclarationParams, Or_textDocument_declaration]
	TextDocumentDefinition              RequestFunc[DefinitionParams, []Location]
	TextDocumentDiagnostic              RequestFunc[DocumentDiagnosticParams, DocumentDiagnosticReport]
	TextDocumentDocumentHighlight       RequestFunc[DocumentHighlightParams, []DocumentHighlight]
	TextDocumentDocumentLink            RequestFunc[DocumentLinkParams, []DocumentLink]
	TextDocumentDocumentSymbol          RequestFunc[DocumentSymbolParams, []any]
	TextDocumentFoldingRange            RequestFunc[FoldingRangeParams, []FoldingRange]
	TextDocumentFormatting              RequestFunc[DocumentFormattingParams, []TextEdit]
	TextDocumentHover                   RequestFunc[HoverParams, Hover]
	TextDocumentImplementation          RequestFunc[ImplementationParams, []Location]
	TextDocumentLinkedEditingRange      RequestFunc[LinkedEditingRangeParams, LinkedEditingRanges]
	TextDocumentMoniker                 RequestFunc[MonikerParams, []Moniker]
	TextDocumentOnTypeFormatting        RequestFunc[DocumentOnTypeFormattingParams, []TextEdit]
	TextDocumentPrepareCallHierarchy    RequestFunc[CallHierarchyPrepareParams, []CallHierarchyItem]
	TextDocumentPrepareRename           RequestFunc[PrepareRenameParams, PrepareRenameResult]
	TextDocumentRangeFormatting         RequestFunc[DocumentRangeFormattingParams, []TextEdit]
	TextDocumentReferences              RequestFunc[ReferenceParams, []Location]
	TextDocumentRename                  RequestFunc[RenameParams, WorkspaceEdit]
	TextDocumentSelectionRange          RequestFunc[SelectionRangeParams, []SelectionRange]
	TextDocumentSemanticTokensFull      RequestFunc[SemanticTokensParams, SemanticTokens]
	TextDocumentSemanticTokensFullDelta RequestFunc[SemanticTokensDeltaParams, any]
	TextDocumentSemanticTokensRange     RequestFunc[SemanticTokensRangeParams, SemanticTokens]
	TextDocumentSignatureHelp           RequestFunc[SignatureHelpParams, SignatureHelp]
	TextDocumentTypeDefinition          RequestFunc[TypeDefinitionParams, []Location]

	WorkspaceDiagnostic     RequestFunc[WorkspaceDiagnosticParams, WorkspaceDiagnosticReport]
	TextDocumentInlineValue RequestFunc[InlineValueParams, []InlineValue]
	TextDocumentInlayHint   RequestFunc[InlayHintParams, []InlayHint]
	InlayHintResolve        RequestFunc[InlayHint, InlayHint]
	WorkspaceSymbolResolve  RequestFunc[SymbolInformation, SymbolInformation]

	common_handler
}
