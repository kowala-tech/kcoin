package discv5

import (
	"bytes"
	"github.com/kowala-tech/kcoin/client/knode/protocol"
)

type Packet struct {
	Prefix     []byte
	Signature  []byte
	Data       []byte
	packetType byte
	topic      []byte
}

// IsDiscoveryPacket returns nil error if a packet is a DiscoveryV5 KCOIN packet
// it's a kind of 'bad and fast' code to make dropping non-kcoin connections as fast as possible
func IsDiscoveryPacket(packetBytes []byte) error {
	packet := newPacket(packetBytes)

	return newChecker(
		packet.isCorrectSize,
		packet.isCorrectPacketType,
		packet.isCorrectPrefix,
		packet.isCorrectPacketData,
	).do()
}

func newPacket(packetBytes []byte) *Packet {
	var (
		prefix    = packetBytes[:versionPrefixSize]
		signature = packetBytes[versionPrefixSize:headSize]
		data      = packetBytes[headSize:]
	)

	packet := &Packet{
		Prefix:     prefix,
		Signature:  signature,
		Data:       data,
		packetType: data[0],
	}
	packet.topic = packet.getPacketTopic()

	return packet
}

func (packet *Packet) getPacketTopic() []byte {
	// further "magic" numbers are from ping, topicRegister, topicQuery types memory model
	// startIndex stands for a position in raw packet where we should find Topic data
	var startIndex int
	switch packet.packetType {
	case pingPacket:
		startIndex = 47
	case topicRegisterPacket:
		startIndex = 4
	case topicQueryPacket:
		startIndex = 3
	}

	endIndex := startIndex + len(protocol.Constants.PrefixBytes)
	if endIndex >= len(packet.Data) {
		return nil
	}

	return packet.Data[startIndex:endIndex]
}

func (packet *Packet) isCorrectSize() error {
	packetLength := len(packet.Prefix) + len(packet.Signature) + len(packet.Data)

	if packetLength <= headSize {
		return errPacketTooSmall
	}
	return nil
}

func (packet *Packet) isCorrectPacketType() error {
	if packet.packetType < pingPacket || packet.packetType > topicNodesPacket {
		return errUnknownPacketType
	}
	return nil
}

func (packet *Packet) isCorrectPrefix() error {
	if !bytes.Equal(packet.Prefix, versionPrefix) {
		return errBadPrefix
	}
	return nil
}

func (packet *Packet) isCorrectPacketData() error {
	return newChecker(
		packet.isCorrectPingPacket,
		packet.isCorrectTopicRegisterPacket,
		packet.isCorrectTopicQueryPacket,
	).do()
}

func (packet *Packet) isCorrectPingPacket() error {
	if packet.packetType != pingPacket {
		return nil
	}
	if packet.isCorrectEmptyPingPacket() {
		return nil
	}
	return packet.isCorrectTopic()
}

func (packet *Packet) isCorrectEmptyPingPacket() bool {
	// a correct ping without a Topic
	return packet.Data[1] == 235
}

func (packet *Packet) isCorrectTopicRegisterPacket() error {
	if packet.packetType != topicRegisterPacket {
		return nil
	}
	return packet.isCorrectTopic()
}

func (packet *Packet) isCorrectTopicQueryPacket() error {
	if packet.packetType != topicQueryPacket {
		return nil
	}
	return packet.isCorrectTopic()
}

func (packet *Packet) isCorrectTopic() error {
	if !bytes.HasPrefix(packet.topic, protocol.Constants.PrefixBytes) {
		return errBadTopic
	}
	return nil
}

type check func() error

type checker struct {
	checks []check
}

func newChecker(checks ...check) *checker {
	return &checker{checks}
}

func (ch *checker) add(checks ...check) *checker {
	ch.checks = append(ch.checks, checks...)
	return ch
}

func (ch *checker) do() error {
	for _, c := range ch.checks {
		if err := c(); err != nil {
			return err
		}
	}
	return nil
}
