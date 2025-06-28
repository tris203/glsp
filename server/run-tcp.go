package server

func (self *Server) RunTCP(address string) error {
	listener, err := self.newNetworkListener("tcp", address)
	if err != nil {
		return err
	}

	log := self.Log.With("address", address)
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

		go self.ServeStream(connection, connectionLog)
	}
}
