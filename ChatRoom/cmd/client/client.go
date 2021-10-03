package client

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/client/handler"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/config"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main(cfg config.Server) {
	os.Mkdir("./client/downloads", 0755)

	connection, err := net.DialTimeout(cfg.Network, cfg.Address, cfg.TimeOut)
	if err != nil {
		logrus.Fatalf("could not establish connection: %s", err.Error())
	}

	client := model.NewClient(connection)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		handler.NewClientHandler(client).Handle()
	}()

	s := <-sig

	logrus.Infof("signal %s received", s)
}

// Register registers client command for chatroom binary.
func Register(root *cobra.Command, cfg config.Server) {
	runServer := &cobra.Command{
		Use:   "client",
		Short: "client for chatroom",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}

	root.AddCommand(runServer)
}
