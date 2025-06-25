package server

import (
	"strings"

	"github.com/tliron/commonlog"
)

type JSONRPCLogger struct {
	log commonlog.Logger
}

// ([jsonrpc2.Logger] interface)
func (l *JSONRPCLogger) Printf(format string, v ...any) {
	l.log.Debugf(strings.TrimSuffix(format, "\n"), v...)
}
