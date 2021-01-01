package cmd

import (
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/PacketSniffer/config"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/PacketSniffer/sniffer"
	"os"
	"time"
)

func Execute() {
	cfg := config.NewPacketSniffer()

	os.Mkdir(cfg.Path, 0755)

	packetSniffer := sniffer.Sniffer{}

	go packetSniffer.Capture(cfg.Path)

	<-time.Tick(10 * time.Second)

	packetSniffer.Finished = true
}
