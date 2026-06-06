package protocol

import (
	"reflect"
)

// Handler names
const (
	HandlerTextDocumentDidOpen                 = "TextDocumentDidOpen"
	HandlerTextDocumentInlineCompletion        = "TextDocumentInlineCompletion"
	HandlerNotebookDocumentDidChange           = "NotebookDocumentDidChange"
	HandlerNotebookDocumentDidClose            = "NotebookDocumentDidClose"
	HandlerNotebookDocumentDidOpen             = "NotebookDocumentDidOpen"
	HandlerNotebookDocumentDidSave             = "NotebookDocumentDidSave"
	HandlerTextDocumentInlineValue             = "TextDocumentInlineValue"
	HandlerTextDocumentInlayHint               = "TextDocumentInlayHint"
	HandlerTextDocumentInlayHintResolve        = "TextDocumentInlayHintResolve"
	HandlerWorkspaceDiagnostic                 = "WorkspaceDiagnostic"
	HandlerWorkspaceSymbolResolve              = "WorkspaceSymbolResolve"
	HandlerTextDocumentDidClose                = "TextDocumentDidClose"
	HandlerTextDocumentDidChange               = "TextDocumentDidChange"
	HandlerTextDocumentWillSave                = "TextDocumentWillSave"
	HandlerTextDocumentWillSaveWaitUntil       = "TextDocumentWillSaveWaitUntil"
	HandlerTextDocumentDidSave                 = "TextDocumentDidSave"
	HandlerTextDocumentCompletion              = "TextDocumentCompletion"
	HandlerTextDocumentHover                   = "TextDocumentHover"
	HandlerTextDocumentSignatureHelp           = "TextDocumentSignatureHelp"
	HandlerTextDocumentDeclaration             = "TextDocumentDeclaration"
	HandlerTextDocumentDefinition              = "TextDocumentDefinition"
	HandlerTextDocumentTypeDefinition          = "TextDocumentTypeDefinition"
	HandlerTextDocumentImplementation          = "TextDocumentImplementation"
	HandlerTextDocumentReferences              = "TextDocumentReferences"
	HandlerTextDocumentDocumentHighlight       = "TextDocumentDocumentHighlight"
	HandlerTextDocumentDocumentSymbol          = "TextDocumentDocumentSymbol"
	HandlerTextDocumentCodeAction              = "TextDocumentCodeAction"
	HandlerTextDocumentCodeLens                = "TextDocumentCodeLens"
	HandlerTextDocumentDocumentLink            = "TextDocumentDocumentLink"
	HandlerTextDocumentDocumentColor           = "TextDocumentDocumentColor"
	HandlerTextDocumentFormatting              = "TextDocumentFormatting"
	HandlerTextDocumentRangeFormatting         = "TextDocumentRangeFormatting"
	HandlerTextDocumentOnTypeFormatting        = "TextDocumentOnTypeFormatting"
	HandlerTextDocumentRename                  = "TextDocumentRename"
	HandlerTextDocumentFoldingRange            = "TextDocumentFoldingRange"
	HandlerWorkspaceExecuteCommand             = "WorkspaceExecuteCommand"
	HandlerTextDocumentSelectionRange          = "TextDocumentSelectionRange"
	HandlerTextDocumentLinkedEditingRange      = "TextDocumentLinkedEditingRange"
	HandlerTextDocumentPrepareCallHierarchy    = "TextDocumentPrepareCallHierarchy"
	HandlerTextDocumentSemanticTokensFull      = "TextDocumentSemanticTokensFull"
	HandlerTextDocumentSemanticTokensRange     = "TextDocumentSemanticTokensRange"
	HandlerTextDocumentSemanticTokensFullDelta = "TextDocumentSemanticTokensFullDelta"
	HandlerTextDocumentMoniker                 = "TextDocumentMoniker"
	HandlerWorkspaceSymbol                     = "WorkspaceSymbol"
	HandlerWorkspaceDidCreateFiles             = "WorkspaceDidCreateFiles"
	HandlerWorkspaceWillCreateFiles            = "WorkspaceWillCreateFiles"
	HandlerWorkspaceDidRenameFiles             = "WorkspaceDidRenameFiles"
	HandlerWorkspaceWillRenameFiles            = "WorkspaceWillRenameFiles"
	HandlerWorkspaceDidDeleteFiles             = "WorkspaceDidDeleteFiles"
	HandlerWorkspaceWillDeleteFiles            = "WorkspaceWillDeleteFiles"
	HandlerTextDocumentDiagnostic              = "TextDocumentDiagnostic"
)

// Capability field names
const (
	CapabilityOpenClose                        = "OpenClose"
	CapabilityChange                           = "Change"
	CapabilityWillSave                         = "WillSave"
	CapabilityWillSaveWaitUntil                = "WillSaveWaitUntil"
	CapabilitySave                             = "Save"
	CapabilityCompletionProvider               = "CompletionProvider"
	CapabilityHoverProvider                    = "HoverProvider"
	CapabilitySignatureHelpProvider            = "SignatureHelpProvider"
	CapabilityDeclarationProvider              = "DeclarationProvider"
	CapabilityDefinitionProvider               = "DefinitionProvider"
	CapabilityTypeDefinitionProvider           = "TypeDefinitionProvider"
	CapabilityImplementationProvider           = "ImplementationProvider"
	CapabilityReferencesProvider               = "ReferencesProvider"
	CapabilityDocumentHighlightProvider        = "DocumentHighlightProvider"
	CapabilityDocumentSymbolProvider           = "DocumentSymbolProvider"
	CapabilityCodeActionProvider               = "CodeActionProvider"
	CapabilityCodeLensProvider                 = "CodeLensProvider"
	CapabilityDocumentLinkProvider             = "DocumentLinkProvider"
	CapabilityColorProvider                    = "ColorProvider"
	CapabilityDocumentFormattingProvider       = "DocumentFormattingProvider"
	CapabilityDocumentRangeFormattingProvider  = "DocumentRangeFormattingProvider"
	CapabilityDocumentOnTypeFormattingProvider = "DocumentOnTypeFormattingProvider"
	CapabilityRenameProvider                   = "RenameProvider"
	CapabilityFoldingRangeProvider             = "FoldingRangeProvider"
	CapabilityExecuteCommandProvider           = "ExecuteCommandProvider"
	CapabilitySelectionRangeProvider           = "SelectionRangeProvider"
	CapabilityLinkedEditingRangeProvider       = "LinkedEditingRangeProvider"
	CapabilityCallHierarchyProvider            = "CallHierarchyProvider"
	CapabilitySemanticTokensProvider           = "SemanticTokensProvider"
	CapabilityMonikerProvider                  = "MonikerProvider"
	CapabilityWorkspaceSymbolProvider          = "WorkspaceSymbolProvider"
	CapabilityDiagnosticProvider               = "DiagnosticProvider"
	CapabilityTextDocumentSync                 = "TextDocumentSync"
	CapabilityFull                             = "Full"
	CapabilityRange                            = "Range"
	CapabilityWorkspace                        = "Workspace"
	CapabilityFileOperations                   = "FileOperations"

	// LSP 3.17+ capabilities
	CapabilityClientRegisterCapability       = "ClientRegisterCapability"
	CapabilityClientUnregisterCapability     = "ClientUnregisterCapability"
	CapabilityWorkspaceDiagnostic            = "WorkspaceDiagnostic"
	CapabilityInlineValueProvider            = "InlineValueProvider"
	CapabilityInlayHintProvider              = "InlayHintProvider"
	CapabilityInlayHintResolveProvider       = "InlayHintResolveProvider"
	CapabilityWorkspaceSymbolResolveProvider = "WorkspaceSymbolResolveProvider"

	// LSP 3.18+ capabilities
	CapabilityInlineCompletionProvider = "InlineCompletionProvider"
	CapabilityNotebookDocumentSync     = "NotebookDocumentSync"
)

// File operation types
const (
	FileOpDidCreate  = "DidCreate"
	FileOpWillCreate = "WillCreate"
	FileOpDidRename  = "DidRename"
	FileOpWillRename = "WillRename"
	FileOpDidDelete  = "DidDelete"
	FileOpWillDelete = "WillDelete"
)

type Handlers interface {
	Handler_316 | Handler_317 | Handler_318
}

func CreateServerCapabilities[H Handlers](handler *H) ServerCapabilities {
	var capabilities ServerCapabilities
	capValue := reflect.ValueOf(&capabilities).Elem()
	handlerValue := reflect.ValueOf(handler)
	// If someone accidentally passes a pointer, unwrap it.
	if handlerValue.Kind() == reflect.Ptr {
		handlerValue = handlerValue.Elem()
	}

	// Text Document Sync handlers
	if hasHandler(handlerValue, HandlerTextDocumentDidOpen) || hasHandler(handlerValue, HandlerTextDocumentDidClose) {
		setTextDocumentSyncOption(capValue, CapabilityOpenClose, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentDidChange) {
		value := Incremental
		setTextDocumentSyncOption(capValue, CapabilityChange, value)
	}

	if hasHandler(handlerValue, HandlerTextDocumentWillSave) {
		setTextDocumentSyncOption(capValue, CapabilityWillSave, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentWillSaveWaitUntil) {
		setTextDocumentSyncOption(capValue, CapabilityWillSaveWaitUntil, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentDidSave) {
		setTextDocumentSyncOption(capValue, CapabilitySave, true)
	}

	// Other capabilities
	if hasHandler(handlerValue, HandlerTextDocumentCompletion) {
		setCapabilityField(capValue, CapabilityCompletionProvider, CompletionOptions{})
	}

	if hasHandler(handlerValue, HandlerTextDocumentHover) {
		setCapabilityField(capValue, CapabilityHoverProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentDeclaration) {
		setCapabilityField(capValue, CapabilityDeclarationProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentDefinition) {
		setCapabilityField(capValue, CapabilityDefinitionProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentTypeDefinition) {
		setCapabilityField(capValue, CapabilityTypeDefinitionProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentImplementation) {
		setCapabilityField(capValue, CapabilityImplementationProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentReferences) {
		setCapabilityField(capValue, CapabilityReferencesProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentDocumentHighlight) {
		setCapabilityField(capValue, CapabilityDocumentHighlightProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentDocumentSymbol) {
		setCapabilityField(capValue, CapabilityDocumentSymbolProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentCodeAction) {
		setCapabilityField(capValue, CapabilityCodeActionProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentDocumentColor) {
		setCapabilityField(capValue, CapabilityColorProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentFormatting) {
		setCapabilityField(capValue, CapabilityDocumentFormattingProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentRangeFormatting) {
		setCapabilityField(capValue, CapabilityDocumentRangeFormattingProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentRename) {
		setCapabilityField(capValue, CapabilityRenameProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentFoldingRange) {
		setCapabilityField(capValue, CapabilityFoldingRangeProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentSelectionRange) {
		setCapabilityField(capValue, CapabilitySelectionRangeProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentLinkedEditingRange) {
		setCapabilityField(capValue, CapabilityLinkedEditingRangeProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentPrepareCallHierarchy) {
		setCapabilityField(capValue, CapabilityCallHierarchyProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentMoniker) {
		setCapabilityField(capValue, CapabilityMonikerProvider, true)
	}

	if hasHandler(handlerValue, HandlerWorkspaceSymbol) {
		setCapabilityField(capValue, CapabilityWorkspaceSymbolProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentSignatureHelp) {
		setCapabilityField(capValue, CapabilitySignatureHelpProvider, SignatureHelpOptions{})
	}

	if hasHandler(handlerValue, HandlerTextDocumentDeclaration) {
		setCapabilityField(capValue, CapabilityDeclarationProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentDefinition) {
		setCapabilityField(capValue, CapabilityDefinitionProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentTypeDefinition) {
		setCapabilityField(capValue, CapabilityTypeDefinitionProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentImplementation) {
		setCapabilityField(capValue, CapabilityImplementationProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentReferences) {
		setCapabilityField(capValue, CapabilityReferencesProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentDocumentHighlight) {
		setCapabilityField(capValue, CapabilityDocumentHighlightProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentDocumentSymbol) {
		setCapabilityField(capValue, CapabilityDocumentSymbolProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentCodeAction) {
		setCapabilityField(capValue, CapabilityCodeActionProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentDocumentColor) {
		setCapabilityField(capValue, CapabilityColorProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentFormatting) {
		setCapabilityField(capValue, CapabilityDocumentFormattingProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentRangeFormatting) {
		setCapabilityField(capValue, CapabilityDocumentRangeFormattingProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentRename) {
		setCapabilityField(capValue, CapabilityRenameProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentFoldingRange) {
		setCapabilityField(capValue, CapabilityFoldingRangeProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentSelectionRange) {
		setCapabilityField(capValue, CapabilitySelectionRangeProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentLinkedEditingRange) {
		setCapabilityField(capValue, CapabilityLinkedEditingRangeProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentPrepareCallHierarchy) {
		setCapabilityField(capValue, CapabilityCallHierarchyProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentMoniker) {
		setCapabilityField(capValue, CapabilityMonikerProvider, true)
	}

	if hasHandler(handlerValue, HandlerWorkspaceSymbol) {
		setCapabilityField(capValue, CapabilityWorkspaceSymbolProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentDefinition) {
		setCapabilityField(capValue, CapabilityDefinitionProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentTypeDefinition) {
		setCapabilityField(capValue, CapabilityTypeDefinitionProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentImplementation) {
		setCapabilityField(capValue, CapabilityImplementationProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentReferences) {
		setCapabilityField(capValue, CapabilityReferencesProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentDocumentHighlight) {
		setCapabilityField(capValue, CapabilityDocumentHighlightProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentDocumentSymbol) {
		setCapabilityField(capValue, CapabilityDocumentSymbolProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentCodeAction) {
		setCapabilityField(capValue, CapabilityCodeActionProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentCodeLens) {
		setCapabilityField(capValue, CapabilityCodeLensProvider, CodeLensOptions{})
	}

	if hasHandler(handlerValue, HandlerTextDocumentDocumentLink) {
		setCapabilityField(capValue, CapabilityDocumentLinkProvider, DocumentLinkOptions{})
	}

	if hasHandler(handlerValue, HandlerTextDocumentDocumentColor) {
		setCapabilityField(capValue, CapabilityColorProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentFormatting) {
		setCapabilityField(capValue, CapabilityDocumentFormattingProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentRangeFormatting) {
		setCapabilityField(capValue, CapabilityDocumentRangeFormattingProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentOnTypeFormatting) {
		setCapabilityField(capValue, CapabilityDocumentOnTypeFormattingProvider, &DocumentOnTypeFormattingOptions{})
	}

	if hasHandler(handlerValue, HandlerTextDocumentRename) {
		setCapabilityField(capValue, CapabilityRenameProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentFoldingRange) {
		setCapabilityField(capValue, CapabilityFoldingRangeProvider, true)
	}

	if hasHandler(handlerValue, HandlerWorkspaceExecuteCommand) {
		setCapabilityField(capValue, CapabilityExecuteCommandProvider, &ExecuteCommandOptions{})
	}

	if hasHandler(handlerValue, HandlerTextDocumentSelectionRange) {
		setCapabilityField(capValue, CapabilitySelectionRangeProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentLinkedEditingRange) {
		setCapabilityField(capValue, CapabilityLinkedEditingRangeProvider, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentPrepareCallHierarchy) {
		setCapabilityField(capValue, CapabilityCallHierarchyProvider, true)
	}

	// Semantic Tokens
	if hasHandler(handlerValue, HandlerTextDocumentSemanticTokensFull) {
		setSemanticTokensOption(capValue, handlerValue)
	}

	if hasHandler(handlerValue, HandlerTextDocumentSemanticTokensRange) {
		setSemanticTokensRangeOption(capValue)
	}

	if hasHandler(handlerValue, HandlerTextDocumentMoniker) {
		setCapabilityField(capValue, CapabilityMonikerProvider, true)
	}

	if hasHandler(handlerValue, HandlerWorkspaceSymbol) {
		setCapabilityField(capValue, CapabilityWorkspaceSymbolProvider, true)
	}

	// File Operations
	if hasHandler(handlerValue, HandlerWorkspaceDidCreateFiles) {
		setFileOperationCapability(capValue, FileOpDidCreate)
	}

	if hasHandler(handlerValue, HandlerWorkspaceWillCreateFiles) {
		setFileOperationCapability(capValue, FileOpWillCreate)
	}

	if hasHandler(handlerValue, HandlerWorkspaceDidRenameFiles) {
		setCapabilityField(capValue, CapabilityRenameProvider, true)
		setFileOperationCapability(capValue, FileOpDidRename)
	}

	if hasHandler(handlerValue, HandlerWorkspaceWillRenameFiles) {
		setCapabilityField(capValue, CapabilityRenameProvider, true)
		setFileOperationCapability(capValue, FileOpWillRename)
	}

	if hasHandler(handlerValue, HandlerWorkspaceDidDeleteFiles) {
		setFileOperationCapability(capValue, FileOpDidDelete)
	}

	if hasHandler(handlerValue, HandlerWorkspaceWillDeleteFiles) {
		setFileOperationCapability(capValue, FileOpWillDelete)
	}

	// LSP 3.17 specific
	if hasHandler(handlerValue, HandlerTextDocumentDiagnostic) {
		diagnosticOptions := DiagnosticOptions{
			InterFileDependencies: true,
			WorkspaceDiagnostics:  false,
		}
		setCapabilityField(capValue, CapabilityDiagnosticProvider, diagnosticOptions)
	}

	// LSP 3.18 specific
	if hasHandler(handlerValue, HandlerTextDocumentInlineCompletion) {
		setCapabilityField(capValue, CapabilityInlineCompletionProvider, true)
	}
	if hasHandler(handlerValue, HandlerNotebookDocumentDidChange) || hasHandler(handlerValue, HandlerNotebookDocumentDidClose) || hasHandler(handlerValue, HandlerNotebookDocumentDidOpen) || hasHandler(handlerValue, HandlerNotebookDocumentDidSave) {
		setCapabilityField(capValue, CapabilityNotebookDocumentSync, true)
	}

	if hasHandler(handlerValue, HandlerTextDocumentInlineValue) {
		setCapabilityField(capValue, CapabilityInlineValueProvider, true)
	}
	if hasHandler(handlerValue, HandlerTextDocumentInlayHint) {
		setCapabilityField(capValue, CapabilityInlayHintProvider, true)
	}
	if hasHandler(handlerValue, HandlerTextDocumentInlayHintResolve) {
		setCapabilityField(capValue, CapabilityInlayHintResolveProvider, true)
	}
	if hasHandler(handlerValue, HandlerWorkspaceDiagnostic) {
		setCapabilityField(capValue, CapabilityWorkspaceDiagnostic, true)
	}
	if hasHandler(handlerValue, HandlerWorkspaceSymbolResolve) {
		setCapabilityField(capValue, CapabilityWorkspaceSymbolResolveProvider, true)
	}

	// ClientRegisterCapability
	// ClientUnregisterCapability

	return capabilities
}

func hasHandler(handlerValue reflect.Value, fieldName string) bool { // fieldName should be a Handler* const
	field := handlerValue.FieldByName(fieldName)
	return field.IsValid() && !field.IsNil()
}

func setCapabilityField(capValue reflect.Value, fieldName string, value any) { // fieldName should be a Capability* const
	field := capValue.FieldByName(fieldName)
	if !field.IsValid() || !field.CanSet() {
		return
	}

	fieldType := field.Type()
	// Check if the field is a pointer to a struct named Or_ServerCapabilities_*
	if fieldType.Kind() == reflect.Ptr && fieldType.Elem().Kind() == reflect.Struct &&
		len(fieldType.Elem().Name()) > len("Or_ServerCapabilities_") &&
		fieldType.Elem().Name()[:len("Or_ServerCapabilities_")] == "Or_ServerCapabilities_" {

		wrapper := reflect.New(fieldType.Elem())
		wrapper.Elem().FieldByName("Value").Set(reflect.ValueOf(value))
		field.Set(wrapper)
		return
	}

	valueOf := reflect.ValueOf(value)
	if valueOf.Type().AssignableTo(fieldType) {
		field.Set(valueOf)
		return
	}

	if fieldType.Kind() == reflect.Ptr && valueOf.Type().AssignableTo(fieldType.Elem()) {
		ptr := reflect.New(fieldType.Elem())
		ptr.Elem().Set(valueOf)
		field.Set(ptr)
		return
	}
}

func setTextDocumentSyncOption(capValue reflect.Value, optionName string, value any) { // optionName should be a Capability* const
	syncField := capValue.FieldByName(CapabilityTextDocumentSync)
	if !syncField.IsValid() || !syncField.CanSet() {
		return
	}

	// Check if it's already a pointer to TextDocumentSyncOptions
	if syncField.Kind() == reflect.Interface {
		if syncField.IsNil() || syncField.Elem().Type().Name() != "TextDocumentSyncOptions" {
			syncOptions := &TextDocumentSyncOptions{}
			syncField.Set(reflect.ValueOf(syncOptions))
		}

		// Get the actual struct
		syncStruct := syncField.Elem().Elem()
		optionField := syncStruct.FieldByName(optionName)
		if optionField.IsValid() && optionField.CanSet() {
			setFieldValue(optionField, value)
		}
	}
}

func setSemanticTokensOption(capValue reflect.Value, handlerValue reflect.Value) {
	providerField := capValue.FieldByName(CapabilitySemanticTokensProvider)
	if !providerField.IsValid() || !providerField.CanSet() {
		return
	}

	if providerField.Kind() == reflect.Interface {
		if providerField.IsNil() || providerField.Elem().Type() != reflect.TypeOf(&SemanticTokensOptions{}) {
			semanticOptions := &SemanticTokensOptions{}
			providerField.Set(reflect.ValueOf(semanticOptions))
		}

		semanticStruct := providerField.Elem().Elem()
		fullField := semanticStruct.FieldByName(CapabilityFull)

		if hasHandler(handlerValue, HandlerTextDocumentSemanticTokensFullDelta) {
			fullField.Set(reflect.ValueOf(&Or_SemanticTokensOptions_full{Value: SemanticTokensFullDelta{Delta: true}}))
		} else {
			fullField.Set(reflect.ValueOf(&Or_SemanticTokensOptions_full{Value: true}))
		}
	}
}

func setSemanticTokensRangeOption(capValue reflect.Value) {
	providerField := capValue.FieldByName(CapabilitySemanticTokensProvider)
	if !providerField.IsValid() || !providerField.CanSet() {
		return
	}

	if providerField.Kind() == reflect.Interface {
		if providerField.IsNil() || providerField.Elem().Type() != reflect.TypeOf(&SemanticTokensOptions{}) {
			semanticOptions := &SemanticTokensOptions{}
			providerField.Set(reflect.ValueOf(semanticOptions))
		}

		semanticStruct := providerField.Elem().Elem()
		rangeField := semanticStruct.FieldByName(CapabilityRange)
		if rangeField.IsValid() && rangeField.CanSet() {
			rangeField.Set(reflect.ValueOf(&Or_SemanticTokensOptions_range{Value: true}))
		}
	}
}

func setFieldValue(field reflect.Value, value any) {
	valueOf := reflect.ValueOf(value)
	fieldType := field.Type()

	if valueOf.Type().AssignableTo(fieldType) {
		field.Set(valueOf)
		return
	}

	if fieldType.Kind() == reflect.Ptr && valueOf.Type().AssignableTo(fieldType.Elem()) {
		ptr := reflect.New(fieldType.Elem())
		ptr.Elem().Set(valueOf)
		field.Set(ptr)
		return
	}

	if fieldType.Kind() == reflect.Ptr && fieldType.Elem().Kind() == reflect.Struct {
		field.Set(reflect.New(fieldType.Elem()))
	}
}

func setFileOperationCapability(capValue reflect.Value, operationType string) { // operationType should be a FileOp* const
	workspaceField := capValue.FieldByName(CapabilityWorkspace)
	if !workspaceField.IsValid() || !workspaceField.CanSet() {
		return
	}

	if workspaceField.IsNil() {
		workspace := &WorkspaceOptions{}
		workspaceField.Set(reflect.ValueOf(workspace))
	}

	fileOpsField := workspaceField.Elem().FieldByName(CapabilityFileOperations)
	if fileOpsField.IsNil() {
		fileOps := &FileOperationOptions{}
		fileOpsField.Set(reflect.ValueOf(fileOps))
	}

	operationField := fileOpsField.Elem().FieldByName(operationType)
	if operationField.IsValid() && operationField.CanSet() {
		registrationOptions := &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
		operationField.Set(reflect.ValueOf(registrationOptions))
	}
}
