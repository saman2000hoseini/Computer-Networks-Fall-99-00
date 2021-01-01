package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/structs"
	"github.com/sirupsen/logrus"
)

type (
	PacketSniffer struct {
		Path string `koanf:"path"`
	}
)

// NewPacketSniffer reads configuration with konaf.
func NewPacketSniffer() PacketSniffer {
	var instance PacketSniffer
	k := koanf.New(".")

	// load default configuration from file
	if err := k.Load(structs.Provider(DefaultPacketSniffer(), "koanf"), nil); err != nil {
		logrus.Fatalf("error loading default: %s", err)
	}

	if err := k.Unmarshal("", &instance); err != nil {
		logrus.Fatalf("error unmarshalling config: %s", err)
	}

	return instance
}
