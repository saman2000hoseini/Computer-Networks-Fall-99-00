package sniffer

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/sirupsen/logrus"
	"os"
	"time"

	"github.com/google/gopacket/pcap"
)

type Sniffer struct {
	Finished bool
}

func (s *Sniffer) Capture(rPath string) {
	fPath := generateName(rPath)

	file, err := os.Create(fPath)
	if err != nil {
		logrus.Errorf("error creating file: %s", err.Error())
		return
	}

	defer file.Close()

	if handle, err := pcap.OpenLive("eth0", 1600, true, pcap.BlockForever); err != nil {
		logrus.Errorf("error openning source: %s", err.Error())
		return
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

		for packet := range packetSource.Packets() {
			if s.Finished {
				break
			}

			file.WriteString(packet.String())
		}
	}

	println(fPath)
}

func generateName(path string) string {
	t := time.Now()

	return fmt.Sprintf("%s/%d%d%d_%d%d%d.txt",
		path, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
