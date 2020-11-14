package cmd

import (
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/client/handler"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/config"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
	"net"
	"os"
)

func Execute() {
	cfg := config.NewServer()

	os.Mkdir("./client/downloads", 0755)

	connection, err := net.DialTimeout(cfg.Network, cfg.Address, cfg.TimeOut)
	if err != nil {
		logrus.Fatalf("could not establish connection: %s", err.Error())
	}

	client := model.NewClient(connection)
	handler.NewClientHandler(client).Handle()
}
