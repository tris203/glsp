package server

import (
	"github.com/tliron/commonlog"
)

func (s *Server) RunTCP(address string) error {
	listener, err := s.newNetworkListener("tcp", address)
	if err != nil {
		return err
	}

	log := commonlog.NewKeyValueLogger(s.Log, "address", address)
	defer commonlog.CallAndLogError((*listener).Close, "listener.Close", log)
	log.Notice("listening for TCP connections")

	var connectionCount uint64

	for {
		connection, err := (*listener).Accept()
		if err != nil {
			return err
		}

		connectionCount++
		connectionLog := commonlog.NewKeyValueLogger(log, "id", connectionCount)

		go s.ServeStream(connection, connectionLog)
	}
}
