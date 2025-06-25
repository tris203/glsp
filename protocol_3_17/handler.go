package protocol

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/tliron/glsp"
	protocol316 "github.com/tliron/glsp/protocol_3_16"
)

type Handler struct {
	protocol316.Handler

	Initialize             InitializeFunc
	TextDocumentDiagnostic TextDocumentDiagnosticFunc

	initialized bool
	lock        sync.Mutex
}

func (h *Handler) Handle(context *glsp.Context) (r any, validMethod bool, validParams bool, err error) {
	if !h.IsInitialized() && (context.Method != protocol316.MethodInitialize) {
		return nil, true, true, errors.New("server not initialized")
	}

	switch context.Method {
	case protocol316.MethodCancelRequest:
		if h.CancelRequest != nil {
			validMethod = true
			var params protocol316.CancelParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = h.CancelRequest(context, &params)
			}
		}

	case protocol316.MethodProgress:
		if h.Progress != nil {
			validMethod = true
			var params protocol316.ProgressParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = h.Progress(context, &params)
			}
		}

	// General Messages

	case MethodInitialize:
		if h.Initialize != nil {
			validMethod = true
			var params InitializeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				if r, err = h.Initialize(context, &params); err == nil {
					h.SetInitialized(true)
				}
			}
		}

	case protocol316.MethodInitialized:
		if h.Initialized != nil {
			validMethod = true
			var params protocol316.InitializedParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = h.Initialized(context, &params)
			}
		}

	case protocol316.MethodShutdown:
		h.SetInitialized(false)
		if h.Shutdown != nil {
			validMethod = true
			validParams = true
			err = h.Shutdown(context)
		}

	case protocol316.MethodExit:
		// Note that the server will close the connection after we handle it here
		if h.Exit != nil {
			validMethod = true
			validParams = true
			err = h.Exit(context)
		}

	case protocol316.MethodLogTrace:
		if h.LogTrace != nil {
			validMethod = true
			var params protocol316.LogTraceParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = h.LogTrace(context, &params)
			}
		}

	case protocol316.MethodSetTrace:
		if h.SetTrace != nil {
			validMethod = true
			var params protocol316.SetTraceParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = h.SetTrace(context, &params)
			}
		}

	// Window

	case protocol316.MethodWindowWorkDoneProgressCancel:
		if h.WindowWorkDoneProgressCancel != nil {
			validMethod = true
			var params protocol316.WorkDoneProgressCancelParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = h.WindowWorkDoneProgressCancel(context, &params)
			}
		}

	// Workspace

	case protocol316.MethodWorkspaceDidChangeWorkspaceFolders:
		if h.WorkspaceDidChangeWorkspaceFolders != nil {
			validMethod = true
			var params protocol316.DidChangeWorkspaceFoldersParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = h.WorkspaceDidChangeWorkspaceFolders(context, &params)
			}
		}

	case protocol316.MethodWorkspaceDidChangeConfiguration:
		if h.WorkspaceDidChangeConfiguration != nil {
			validMethod = true
			var params protocol316.DidChangeConfigurationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = h.WorkspaceDidChangeConfiguration(context, &params)
			}
		}

	case protocol316.MethodWorkspaceDidChangeWatchedFiles:
		if h.WorkspaceDidChangeWatchedFiles != nil {
			validMethod = true
			var params protocol316.DidChangeWatchedFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = h.WorkspaceDidChangeWatchedFiles(context, &params)
			}
		}

	case protocol316.MethodWorkspaceSymbol:
		if h.WorkspaceSymbol != nil {
			validMethod = true
			var params protocol316.WorkspaceSymbolParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.WorkspaceSymbol(context, &params)
			}
		}

	case protocol316.MethodWorkspaceExecuteCommand:
		if h.WorkspaceExecuteCommand != nil {
			validMethod = true
			var params protocol316.ExecuteCommandParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.WorkspaceExecuteCommand(context, &params)
			}
		}

	case protocol316.MethodWorkspaceWillCreateFiles:
		if h.WorkspaceWillCreateFiles != nil {
			validMethod = true
			var params protocol316.CreateFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.WorkspaceWillCreateFiles(context, &params)
			}
		}

	case protocol316.MethodWorkspaceDidCreateFiles:
		if h.WorkspaceDidCreateFiles != nil {
			validMethod = true
			var params protocol316.CreateFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = h.WorkspaceDidCreateFiles(context, &params)
			}
		}

	case protocol316.MethodWorkspaceWillRenameFiles:
		if h.WorkspaceWillRenameFiles != nil {
			validMethod = true
			var params protocol316.RenameFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.WorkspaceWillRenameFiles(context, &params)
			}
		}

	case protocol316.MethodWorkspaceDidRenameFiles:
		if h.WorkspaceDidRenameFiles != nil {
			validMethod = true
			var params protocol316.RenameFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = h.WorkspaceDidRenameFiles(context, &params)
			}
		}

	case protocol316.MethodWorkspaceWillDeleteFiles:
		if h.WorkspaceWillDeleteFiles != nil {
			validMethod = true
			var params protocol316.DeleteFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.WorkspaceWillDeleteFiles(context, &params)
			}
		}

	case protocol316.MethodWorkspaceDidDeleteFiles:
		if h.WorkspaceDidDeleteFiles != nil {
			validMethod = true
			var params protocol316.DeleteFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = h.WorkspaceDidDeleteFiles(context, &params)
			}
		}

	// Text Document Synchronization

	case protocol316.MethodTextDocumentDidOpen:
		if h.TextDocumentDidOpen != nil {
			validMethod = true
			var params protocol316.DidOpenTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = h.TextDocumentDidOpen(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDidChange:
		if h.TextDocumentDidChange != nil {
			validMethod = true
			var params protocol316.DidChangeTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = h.TextDocumentDidChange(context, &params)
			}
		}

	case protocol316.MethodTextDocumentWillSave:
		if h.TextDocumentWillSave != nil {
			validMethod = true
			var params protocol316.WillSaveTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = h.TextDocumentWillSave(context, &params)
			}
		}

	case protocol316.MethodTextDocumentWillSaveWaitUntil:
		if h.TextDocumentWillSaveWaitUntil != nil {
			validMethod = true
			var params protocol316.WillSaveTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentWillSaveWaitUntil(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDidSave:
		if h.TextDocumentDidSave != nil {
			validMethod = true
			var params protocol316.DidSaveTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = h.TextDocumentDidSave(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDidClose:
		if h.TextDocumentDidClose != nil {
			validMethod = true
			var params protocol316.DidCloseTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = h.TextDocumentDidClose(context, &params)
			}
		}

	// Language Features

	case protocol316.MethodTextDocumentCompletion:
		if h.TextDocumentCompletion != nil {
			validMethod = true
			var params protocol316.CompletionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentCompletion(context, &params)
			}
		}

	case protocol316.MethodCompletionItemResolve:
		if h.CompletionItemResolve != nil {
			validMethod = true
			var params protocol316.CompletionItem
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.CompletionItemResolve(context, &params)
			}
		}

	case protocol316.MethodTextDocumentHover:
		if h.TextDocumentHover != nil {
			validMethod = true
			var params protocol316.HoverParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentHover(context, &params)
			}
		}

	case protocol316.MethodTextDocumentSignatureHelp:
		if h.TextDocumentSignatureHelp != nil {
			validMethod = true
			var params protocol316.SignatureHelpParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentSignatureHelp(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDeclaration:
		if h.TextDocumentDeclaration != nil {
			validMethod = true
			var params protocol316.DeclarationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentDeclaration(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDefinition:
		if h.TextDocumentDefinition != nil {
			validMethod = true
			var params protocol316.DefinitionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentDefinition(context, &params)
			}
		}

	case protocol316.MethodTextDocumentTypeDefinition:
		if h.TextDocumentTypeDefinition != nil {
			validMethod = true
			var params protocol316.TypeDefinitionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentTypeDefinition(context, &params)
			}
		}

	case protocol316.MethodTextDocumentImplementation:
		if h.TextDocumentImplementation != nil {
			validMethod = true
			var params protocol316.ImplementationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentImplementation(context, &params)
			}
		}

	case protocol316.MethodTextDocumentReferences:
		if h.TextDocumentReferences != nil {
			validMethod = true
			var params protocol316.ReferenceParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentReferences(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDocumentHighlight:
		if h.TextDocumentDocumentHighlight != nil {
			validMethod = true
			var params protocol316.DocumentHighlightParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentDocumentHighlight(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDocumentSymbol:
		if h.TextDocumentDocumentSymbol != nil {
			validMethod = true
			var params protocol316.DocumentSymbolParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentDocumentSymbol(context, &params)
			}
		}

	case protocol316.MethodTextDocumentCodeAction:
		if h.TextDocumentCodeAction != nil {
			validMethod = true
			var params protocol316.CodeActionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentCodeAction(context, &params)
			}
		}

	case protocol316.MethodCodeActionResolve:
		if h.CodeActionResolve != nil {
			validMethod = true
			var params protocol316.CodeAction
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.CodeActionResolve(context, &params)
			}
		}

	case protocol316.MethodTextDocumentCodeLens:
		if h.TextDocumentCodeLens != nil {
			validMethod = true
			var params protocol316.CodeLensParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentCodeLens(context, &params)
			}
		}

	case protocol316.MethodCodeLensResolve:
		if h.TextDocumentDidClose != nil {
			validMethod = true
			var params protocol316.CodeLens
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.CodeLensResolve(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDocumentLink:
		if h.TextDocumentDocumentLink != nil {
			validMethod = true
			var params protocol316.DocumentLinkParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentDocumentLink(context, &params)
			}
		}

	case protocol316.MethodDocumentLinkResolve:
		if h.DocumentLinkResolve != nil {
			validMethod = true
			var params protocol316.DocumentLink
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.DocumentLinkResolve(context, &params)
			}
		}

	case protocol316.MethodTextDocumentColor:
		if h.TextDocumentColor != nil {
			validMethod = true
			var params protocol316.DocumentColorParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentColor(context, &params)
			}
		}

	case protocol316.MethodTextDocumentColorPresentation:
		if h.TextDocumentColorPresentation != nil {
			validMethod = true
			var params protocol316.ColorPresentationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentColorPresentation(context, &params)
			}
		}

	case protocol316.MethodTextDocumentFormatting:
		if h.TextDocumentFormatting != nil {
			validMethod = true
			var params protocol316.DocumentFormattingParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentFormatting(context, &params)
			}
		}

	case protocol316.MethodTextDocumentRangeFormatting:
		if h.TextDocumentRangeFormatting != nil {
			validMethod = true
			var params protocol316.DocumentRangeFormattingParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentRangeFormatting(context, &params)
			}
		}

	case protocol316.MethodTextDocumentOnTypeFormatting:
		if h.TextDocumentOnTypeFormatting != nil {
			validMethod = true
			var params protocol316.DocumentOnTypeFormattingParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentOnTypeFormatting(context, &params)
			}
		}

	case protocol316.MethodTextDocumentRename:
		if h.TextDocumentRename != nil {
			validMethod = true
			var params protocol316.RenameParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentRename(context, &params)
			}
		}

	case protocol316.MethodTextDocumentPrepareRename:
		if h.TextDocumentPrepareRename != nil {
			validMethod = true
			var params protocol316.PrepareRenameParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentPrepareRename(context, &params)
			}
		}

	case protocol316.MethodTextDocumentFoldingRange:
		if h.TextDocumentFoldingRange != nil {
			validMethod = true
			var params protocol316.FoldingRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentFoldingRange(context, &params)
			}
		}

	case protocol316.MethodTextDocumentSelectionRange:
		if h.TextDocumentSelectionRange != nil {
			validMethod = true
			var params protocol316.SelectionRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentSelectionRange(context, &params)
			}
		}

	case protocol316.MethodTextDocumentPrepareCallHierarchy:
		if h.TextDocumentPrepareCallHierarchy != nil {
			validMethod = true
			var params protocol316.CallHierarchyPrepareParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentPrepareCallHierarchy(context, &params)
			}
		}

	case protocol316.MethodCallHierarchyIncomingCalls:
		if h.CallHierarchyIncomingCalls != nil {
			validMethod = true
			var params protocol316.CallHierarchyIncomingCallsParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.CallHierarchyIncomingCalls(context, &params)
			}
		}

	case protocol316.MethodCallHierarchyOutgoingCalls:
		if h.CallHierarchyOutgoingCalls != nil {
			validMethod = true
			var params protocol316.CallHierarchyOutgoingCallsParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.CallHierarchyOutgoingCalls(context, &params)
			}
		}

	case protocol316.MethodTextDocumentSemanticTokensFull:
		if h.TextDocumentSemanticTokensFull != nil {
			validMethod = true
			var params protocol316.SemanticTokensParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentSemanticTokensFull(context, &params)
			}
		}

	case protocol316.MethodTextDocumentSemanticTokensFullDelta:
		if h.TextDocumentSemanticTokensFullDelta != nil {
			validMethod = true
			var params protocol316.SemanticTokensDeltaParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentSemanticTokensFullDelta(context, &params)
			}
		}

	case protocol316.MethodTextDocumentSemanticTokensRange:
		if h.TextDocumentSemanticTokensRange != nil {
			validMethod = true
			var params protocol316.SemanticTokensRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentSemanticTokensRange(context, &params)
			}
		}

	case protocol316.MethodWorkspaceSemanticTokensRefresh:
		if h.WorkspaceSemanticTokensRefresh != nil {
			validMethod = true
			validParams = true
			err = h.WorkspaceSemanticTokensRefresh(context)
		}

	case protocol316.MethodTextDocumentLinkedEditingRange:
		if h.TextDocumentLinkedEditingRange != nil {
			validMethod = true
			var params protocol316.LinkedEditingRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentLinkedEditingRange(context, &params)
			}
		}

	case protocol316.MethodTextDocumentMoniker:
		if h.TextDocumentMoniker != nil {
			validMethod = true
			var params protocol316.MonikerParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentMoniker(context, &params)
			}
		}
	case MethodTextDocumentDiagnostic:
		if h.TextDocumentDiagnostic != nil {
			validMethod = true
			var params DocumentDiagnosticParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = h.TextDocumentDiagnostic(context, &params)
			}
		}

	default:
		if h.CustomRequest != nil {
			if handler, ok := h.CustomRequest[context.Method]; ok && (handler.Func != nil) {
				validMethod = true
				if err = json.Unmarshal(context.Params, &handler.Params); err == nil {
					validParams = true
					r, err = handler.Func(context, handler.Params)
				}
			}
		}
	}

	return

}

func (h *Handler) IsInitialized() bool {
	h.lock.Lock()
	defer h.lock.Unlock()
	return h.initialized
}

func (h *Handler) SetInitialized(initialized bool) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.initialized = initialized
}

func (h *Handler) CreateServerCapabilities() ServerCapabilities {
	var capabilities ServerCapabilities

	if (h.TextDocumentDidOpen != nil) || (h.TextDocumentDidClose != nil) {
		if _, ok := capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol316.TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions).OpenClose = &protocol316.True
	}

	if h.TextDocumentDidChange != nil {
		if _, ok := capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol316.TextDocumentSyncOptions{}
		}
		// This can be overriden to TextDocumentSyncKindFull
		value := protocol316.TextDocumentSyncKindIncremental
		capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions).Change = &value
	}

	if h.TextDocumentWillSave != nil {
		if _, ok := capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol316.TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions).WillSave = &protocol316.True
	}

	if h.TextDocumentWillSaveWaitUntil != nil {
		if _, ok := capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol316.TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions).WillSaveWaitUntil = &protocol316.True
	}

	if h.TextDocumentDidSave != nil {
		if _, ok := capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol316.TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions).Save = &protocol316.True
	}

	if h.TextDocumentCompletion != nil {
		capabilities.CompletionProvider = &protocol316.CompletionOptions{}
	}

	if h.TextDocumentHover != nil {
		capabilities.HoverProvider = true
	}

	if h.TextDocumentSignatureHelp != nil {
		capabilities.SignatureHelpProvider = &protocol316.SignatureHelpOptions{}
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
		capabilities.CodeLensProvider = &protocol316.CodeLensOptions{}
	}

	if h.TextDocumentDocumentLink != nil {
		capabilities.DocumentLinkProvider = &protocol316.DocumentLinkOptions{}
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
		capabilities.DocumentOnTypeFormattingProvider = &protocol316.DocumentOnTypeFormattingOptions{}
	}

	if h.TextDocumentRename != nil {
		capabilities.RenameProvider = true
	}

	if h.TextDocumentFoldingRange != nil {
		capabilities.FoldingRangeProvider = true
	}

	if h.WorkspaceExecuteCommand != nil {
		capabilities.ExecuteCommandProvider = &protocol316.ExecuteCommandOptions{}
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
		if _, ok := capabilities.SemanticTokensProvider.(*protocol316.SemanticTokensOptions); !ok {
			capabilities.SemanticTokensProvider = &protocol316.SemanticTokensOptions{}
		}
		if h.TextDocumentSemanticTokensFullDelta != nil {
			capabilities.SemanticTokensProvider.(*protocol316.SemanticTokensOptions).Full = &protocol316.SemanticDelta{}
			capabilities.SemanticTokensProvider.(*protocol316.SemanticTokensOptions).Full.(*protocol316.SemanticDelta).Delta = &protocol316.True
		} else {
			capabilities.SemanticTokensProvider.(*protocol316.SemanticTokensOptions).Full = true
		}
	}

	if h.TextDocumentSemanticTokensRange != nil {
		if _, ok := capabilities.SemanticTokensProvider.(*protocol316.SemanticTokensOptions); !ok {
			capabilities.SemanticTokensProvider = &protocol316.SemanticTokensOptions{}
		}
		capabilities.SemanticTokensProvider.(*protocol316.SemanticTokensOptions).Range = true
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
			capabilities.Workspace = &protocol316.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol316.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidCreate = &protocol316.FileOperationRegistrationOptions{
			Filters: []protocol316.FileOperationFilter{},
		}
	}

	if h.WorkspaceWillCreateFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol316.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol316.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillCreate = &protocol316.FileOperationRegistrationOptions{
			Filters: []protocol316.FileOperationFilter{},
		}
	}

	if h.WorkspaceDidRenameFiles != nil {
		capabilities.RenameProvider = true
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol316.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol316.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidRename = &protocol316.FileOperationRegistrationOptions{
			Filters: []protocol316.FileOperationFilter{},
		}
	}

	if h.WorkspaceWillRenameFiles != nil {
		capabilities.RenameProvider = true
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol316.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol316.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillRename = &protocol316.FileOperationRegistrationOptions{
			Filters: []protocol316.FileOperationFilter{},
		}
	}

	if h.WorkspaceDidDeleteFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol316.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol316.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidDelete = &protocol316.FileOperationRegistrationOptions{
			Filters: []protocol316.FileOperationFilter{},
		}
	}

	if h.WorkspaceWillDeleteFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol316.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol316.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillDelete = &protocol316.FileOperationRegistrationOptions{
			Filters: []protocol316.FileOperationFilter{},
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
