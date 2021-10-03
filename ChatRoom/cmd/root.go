package cmd

import (
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/cmd/client"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/cmd/server"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/config"
	"github.com/spf13/cobra"
)

// NewRootCommand creates a new chat room root command.
func NewRootCommand() *cobra.Command {
	var root = &cobra.Command{
		Use: "chatroom",
	}

	cfgServer := config.NewServer()

	server.Register(root, cfgServer)
	client.Register(root, cfgServer)

	return root
}
