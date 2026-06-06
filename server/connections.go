package server

import (
	"context"
	"io"

	"github.com/gorilla/websocket"
	"github.com/sourcegraph/jsonrpc2"
	wsjsonrpc2 "github.com/sourcegraph/jsonrpc2/websocket"
)

func (s *Server) newStreamConnection(stream io.ReadWriteCloser) *jsonrpc2.Conn {
	handler := s.newHandler()
	connectionOptions := s.newConnectionOptions()

	context, cancel := context.WithTimeout(context.Background(), s.StreamTimeout)
	defer cancel()

	return jsonrpc2.NewConn(context, jsonrpc2.NewBufferedStream(stream, jsonrpc2.VSCodeObjectCodec{}), handler, connectionOptions...)
}

func (s *Server) newWebSocketConnection(socket *websocket.Conn) *jsonrpc2.Conn {
	handler := s.newHandler()
	connectionOptions := s.newConnectionOptions()

	context, cancel := context.WithTimeout(context.Background(), s.WebSocketTimeout)
	defer cancel()

	return jsonrpc2.NewConn(context, wsjsonrpc2.NewObjectStream(socket), handler, connectionOptions...)
}

func (s *Server) newConnectionOptions() []jsonrpc2.ConnOpt {
	if s.Debug {
		log := s.Log.With("scope", "rpc")
		return []jsonrpc2.ConnOpt{jsonrpc2.LogMessages(&JSONRPCLogger{log})}
	} else {
		return nil
	}
}
