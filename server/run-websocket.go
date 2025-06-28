package server

import (
	"net/http"
	"sync/atomic"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

func (self *Server) RunWebSocket(address string) error {
	mux := http.NewServeMux()
	upgrader := websocket.Upgrader{CheckOrigin: func(request *http.Request) bool { return true }}

	var connectionCount uint64

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		connection, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			self.Log.Warn("error upgrading HTTP to web socket: %s", "error", err.Error())
			http.Error(writer, errors.Wrap(err, "could not upgrade to web socket").Error(), http.StatusBadRequest)
			return
		}

		log := self.Log.With("id", atomic.AddUint64(&connectionCount, 1))
		defer func() {
			err := connection.Close()
			if err != nil {
				log.Error("error closing web socket connection", "error", err.Error())
			}
		}()
		self.ServeWebSocket(connection, log)
	})

	listener, err := self.newNetworkListener("tcp", address)
	if err != nil {
		return err
	}

	server := http.Server{
		Handler:      http.TimeoutHandler(mux, self.Timeout, ""),
		ReadTimeout:  self.ReadTimeout,
		WriteTimeout: self.WriteTimeout,
	}

	self.Log.Info("listening for web socket connections", "address", address)
	err = server.Serve(*listener)
	return errors.Wrap(err, "WebSocket")
}
