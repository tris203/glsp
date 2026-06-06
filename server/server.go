package server

import (
	"log/slog"
	"time"

	"github.com/tliron/glsp"
)

var DefaultTimeout = time.Minute

//
// Server
//

type Server struct {
	Handler glsp.RPCHandler
	Debug   bool

	Log              *slog.Logger
	Timeout          time.Duration
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	StreamTimeout    time.Duration
	WebSocketTimeout time.Duration
}

func NewServer(handler glsp.RPCHandler, logger *slog.Logger, debug bool) *Server {
	return &Server{
		Handler:          handler,
		Debug:            debug,
		Log:              logger,
		Timeout:          DefaultTimeout,
		ReadTimeout:      DefaultTimeout,
		WriteTimeout:     DefaultTimeout,
		StreamTimeout:    DefaultTimeout,
		WebSocketTimeout: DefaultTimeout,
	}
}
