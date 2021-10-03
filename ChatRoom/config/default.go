package config

import "time"

// DefaultServer return default configuration
//nolint:gomnd
func DefaultServer() Server {
	return Server{
		Network: "tcp",
		Address: ":65432",
		TimeOut: 10 * time.Second,
	}
}
