package config

import (
	"github.com/kowala-tech/kUSD/kusd"
	"github.com/kowala-tech/kUSD/node"
)

// Config defines the top level configuration for a Kowala node
type Config struct {
	Node   *node.Config `mapstructure:"node"`
	Kowala *kusd.Config `mapstructure:"kusd"`
}

// DefaultConfig returns a default configuration for a Kowala node
func DefaultConfig() *Config {
	return &Config{
		Node:   node.DefaultConfig,
		Kowala: kusd.DefaultConfig,
	}
}
