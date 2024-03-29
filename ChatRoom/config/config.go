package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/structs"
	"github.com/sirupsen/logrus"
	"time"
)

type (
	Server struct {
		Network string        `koanf:"network"`
		Address string        `koanf:"address"`
		TimeOut time.Duration `koanf:"time-out"`
	}
)

// NewServer reads configuration with konaf.
func NewServer() Server {
	var instance Server
	k := koanf.New(".")

	// load default configuration from file
	if err := k.Load(structs.Provider(DefaultServer(), "koanf"), nil); err != nil {
		logrus.Fatalf("error loading default: %s", err)
	}

	if err := k.Unmarshal("", &instance); err != nil {
		logrus.Fatalf("error unmarshalling config: %s", err)
	}

	return instance
}
