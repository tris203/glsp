package server

func (s *Server) RunTCP(address string) error {
	listener, err := s.newNetworkListener("tcp", address)
	if err != nil {
		return err
	}

	log := s.Log.With("address", address)
	defer func() {
		err := (*listener).Close()
		if err != nil {
			log.Error("error closing tcp connection", "error", err.Error())
		}
	}()

	log.Info("listening for TCP connections", "address", address)

	var connectionCount uint64

	for {
		connection, err := (*listener).Accept()
		if err != nil {
			return err
		}

		connectionCount++
		connectionLog := log.With("id", connectionCount)

		go s.ServeStream(connection, connectionLog)
	}
}
