package server

//Config represents the configuration accepted by the server.
type Config struct {
	listenPort uint16
}

// DefaultConfig returns a config struct with the default
// configuration options.
func DefaultConfig() *Config {
	return &Config{
		listenPort: 8080,
	}
}
