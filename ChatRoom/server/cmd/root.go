package cmd

import (
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/config"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/db"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/handler"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
	"net"
)

func Execute() {
	myDB, err := db.FirstSetup()
	if err != nil {
		logrus.Fatalf("failed to setup db: %s", err.Error())
	}

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

		clientHandler := handler.NewClientHandler(model.NewClient(connection), myDB)
		go clientHandler.StartListening()
	}
}
