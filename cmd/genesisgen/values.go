package main

import (
	"errors"
)

var (
	MainNetwork  = "main"
	TestNetwork  = "test"
	OtherNetwork = "other"

	AvailableNetworks = map[string]bool{
		MainNetwork : true,
		TestNetwork : true,
		OtherNetwork: true,
	}

	ErrInvalidNetwork = errors.New("invalid network, use main, test or other")
)

//NewNetwork checks and returns a string that represents a network. Maybe it is a good idea
//to have some value object that represents in the entire codebase the available networks.
func NewNetwork(network string) (string, error) {
	if !AvailableNetworks[network] {
		return "", ErrInvalidNetwork
	}

	return network, nil
}
