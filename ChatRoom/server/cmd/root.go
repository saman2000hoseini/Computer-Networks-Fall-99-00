package cmd

import (
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/config"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/handler"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"net"
)

func Execute() {
	cfg := config.NewServer()

	server, err := net.Listen(cfg.Network, cfg.Address)
	if err != nil {
		panic(err)
	}
	for {
		connection, err := server.Accept()
		if err != nil {
			panic(err)
		}

		clientHandler := handler.NewClientHandler(model.NewClient(connection))
		go clientHandler.StartListening()
	}
}
