package cmd

import (
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/PacketSniffer/config"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/PacketSniffer/sniffer"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/PacketSniffer/ui"
	"os"
)

func Execute() {
	cfg := config.NewPacketSniffer()

	os.Mkdir(cfg.Path, 0755)

	packetSniffer := new(sniffer.Sniffer)
	cli := &ui.CLI{PacketSniffer: packetSniffer, Path: cfg.Path}

	cli.Start()
}
