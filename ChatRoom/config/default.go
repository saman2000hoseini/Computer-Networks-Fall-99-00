package config

// Default return default configuration
//nolint:gomnd
func DefaultServer() Server {
	return Server{
		Network: "tcp",
		Address: ":65432",
	}
}
