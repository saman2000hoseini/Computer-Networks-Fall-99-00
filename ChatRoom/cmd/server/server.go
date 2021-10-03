package server

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/config"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/db"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/handler"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main(cfg config.Server) {
	myDB, err := db.FirstSetup()
	if err != nil {
		logrus.Fatalf("failed to setup db: %s", err.Error())
	}

	os.Mkdir("./server/storage", 0755)

	server, err := net.Listen(cfg.Network, cfg.Address)
	logrus.Info("Server has started")

	clientHandler := handler.NewClientHandler(myDB)
	if err != nil {
		panic(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			connection, err := server.Accept()
			if err != nil {
				panic(err)
			}
			logrus.Info("Someone joined")

			client := model.NewClient(connection)
			go clientHandler.StartListening(client)
			go clientHandler.Respond(client)
			go clientHandler.HandleRequest(client)
		}
	}()

	logrus.Info("chatroom server started!")

	s := <-sig

	logrus.Infof("signal %s received", s)
}

// Register registers server command for chatroom binary.
func Register(root *cobra.Command, cfg config.Server) {
	runServer := &cobra.Command{
		Use:   "server",
		Short: "server for chatroom",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}

	root.AddCommand(runServer)
}
