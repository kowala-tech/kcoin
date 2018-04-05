package kcoin

import (
	"errors"
)

const (
	DefaultSmartContractsOwner = "0x259be75d96876f2ada3d202722523e9cd4dd917d"

	MainNetwork  = "main"
	TestNetwork  = "test"
	OtherNetwork = "other"

	TendermintConsensus = "tendermint"
)

var (
	AvailableNetworks = map[string]bool{
		MainNetwork : true,
		TestNetwork : true,
		OtherNetwork: true,
	}

	AvailableConsensusEngine = map[string]bool {
		TendermintConsensus: true,
	}

	ErrInvalidNetwork         = errors.New("invalid Network, use main, test or other")
	ErrInvalidConsensusEngine = errors.New("invalid consensus engine")
)

//NewNetwork checks and returns a string that represents a Network. Maybe it is a good idea
//to have some value object that represents in the entire codebase the available networks.
func NewNetwork(network string) (string, error) {
	if !AvailableNetworks[network] {
		return "", ErrInvalidNetwork
	}

	return network, nil
}

func NewConsensusEngine(consensus string) (string, error) {
	if !AvailableConsensusEngine[consensus]	 {
		return "", ErrInvalidConsensusEngine
	}

	return consensus, nil
}