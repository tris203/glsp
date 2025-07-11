// Package shared provides shared constants and helpers for code generation tools.
package shared

const (
	// VSCodeRepo is the URL of the VSCode LSP protocol repository.
	VSCodeRepo = "https://github.com/microsoft/vscode-languageserver-node"
)

// LSPGitRef is the branch or tag in VSCodeRepo used for protocol versioning.
var LSPGitRef = "release/protocol/3.17.6-next.9"
