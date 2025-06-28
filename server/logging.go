package server

import (
	"log/slog"
	"strings"
)

type JSONRPCLogger struct {
	log *slog.Logger
}

// ([jsonrpc2.Logger] interface)
func (self *JSONRPCLogger) Printf(format string, v ...any) {
	self.log.Debug(strings.TrimSuffix(format, "\n"), v...)
}
