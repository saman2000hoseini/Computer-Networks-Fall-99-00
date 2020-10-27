package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/sirupsen/logrus"
)

type (
	Server struct {
		Network string `koanf:"network"`
		Address string `koanf:"address"`
	}
)

// New reads configuration with konaf.
func NewServer() Server {
	var instance Server
	k := koanf.New(".")

	// load default configuration from file
	if err := k.Load(structs.Provider(DefaultServer(), "koanf"), nil); err != nil {
		logrus.Fatalf("error loading default: %s", err)
	}

	// load configuration from file
	if err := k.Load(file.Provider("config.yml"), yaml.Parser()); err != nil {
		logrus.Errorf("error loading config.yml: %s", err)
	}

	if err := k.Unmarshal("", &instance); err != nil {
		logrus.Fatalf("error unmarshalling config: %s", err)
	}

	return instance
}
