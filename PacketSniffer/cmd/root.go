package cmd

import (
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/PacketSniffer/config"
	"os"
)

func Execute() {
	cfg := config.NewPacketSniffer()

	os.Mkdir("./"+cfg.Path, 0755)
}
