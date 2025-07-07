_This is an early release. Some features are not yet fully implemented._

# GLSP

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Reference](https://pkg.go.dev/badge/github.com/tliron/glsp.svg)](https://pkg.go.dev/github.com/tliron/kutglspil)
[![Go Report Card](https://goreportcard.com/badge/github.com/tliron/glsp)](https://goreportcard.com/report/github.com/tliron/glsp)

[Language Server Protocol](https://microsoft.github.io/language-server-protocol/) SDK for Go.

It enables you to more easily implement language servers by writing them in Go. GLSP contains:

1. all the message structures for easy serialization,
2. a handler for all client methods, and
3. a ready-to-run JSON-RPC 2.0 server supporting stdio, TCP, WebSockets, and Node.js IPC.

All you need to do, then, is provide the features for the language you want to support.

Projects using GLSP:

- [Puccini TOSCA Language Server](https://github.com/tliron/puccini-language-server)
- [zk](https://github.com/mickael-menu/zk)

## References

- [go-lsp](https://github.com/sourcegraph/go-lsp) is another implementation with reduced coverage of the protocol

## Minimal Example

```go
package main

import (
	"fmt"
	"log/slog"

	"github.com/tliron/glsp"
	"github.com/tliron/glsp/protocol"
	"github.com/tliron/glsp/server"
)

const lsName = "my language"

var (
	version string = "0.0.1"
	handler protocol.Handler_316
)

func main() {
	handler = protocol.Handler_316{
		Initialize:  initialize,
		Initialized: initialized,
		Shutdown:    shutdown,
		SetTrace:    setTrace,
	}

	protocol.AddCustomRequest("test/test", TestHandler)
	protocol.AddCustomNotification("test/noti", TestNotificationHandler)

	logger := slog.Default()
	server := server.NewServer(&handler, logger, false)

	server.RunStdio()
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (protocol.InitializeResult, error) {
	capabilities := protocol.CreateServerCapabilities(&handler)

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.ServerInfo{
			Name:    lsName,
			Version: version,
		},
	}, nil
}

func initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	return nil
}

func shutdown(context *glsp.Context) error {
	protocol.SetTraceValue(protocol.Off)
	return nil
}

type TestParams struct {
	From_Client string `json:"from_client"`
}

type TestResult struct {
	From_Server string `json:"out"`
}

func TestHandler(context *glsp.Context, params *TestParams) (TestResult, error) {
	fmt.Printf("Parameters: %v\n", params.From_Client)
	return TestResult{From_Server: "Hello From Server"}, nil
}

func TestNotificationHandler(context *glsp.Context, params *TestParams) error {
	fmt.Println("Test Notification Handler")
	fmt.Printf("Parameters: %v\n", params.From_Client)
	return nil
}

func setTrace(context *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}
```
