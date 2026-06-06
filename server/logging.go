package server

import (
	"log/slog"
	"strings"
)

type JSONRPCLogger struct {
	log *slog.Logger
}

// ([jsonrpc2.Logger] interface)
func (s *JSONRPCLogger) Printf(format string, v ...any) {
	s.log.Debug(strings.TrimSuffix(format, "\n"), v...)
}
