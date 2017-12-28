package config

import (
	"github.com/kowala-tech/kUSD/node"
	"github.com/kowala-tech/kUSD/stats"
)

// Config defines the top level configuration for a Kowala node
type Config struct {
	Node      *node.Config      `mapstructure:"node"`
	Stats     *stats.Config     `mapstructure:"stats"`
	Consensus *consensus.Config `mapstructure:"consensus"`
}

// DefaultConfig returns a default configuration for a Kowala node
func DefaultConfig() *Config {
	return &Config{
		Node:      node.DefaultConfig,
		Stats:     stats.DefaultConfig,
		Consensus: consensus.DefaultConfig,
	}
}
