package protocol

import (
	"core/byteutil"
	"encoding/binary"
	"errors"
	"core/network/gossip/protocol/interfaces"
	"core/network/gossip/protocol/version1"
	"strconv"
)

const (
	PROTOCOL_VERSION_PACKET_SIZE = 4
)

type ProtocolVersionState struct {
	buffer []byte
	offset int
}

func NewProtocolVersionState() *ProtocolVersionState {
	this := &ProtocolVersionState{[]byte{0, 0, 0, 0}, 0}

	return this
}

func (this *ProtocolVersionState) Consume(protocol interfaces.Protocol, data []byte, offset int, length int) (int, error) {
	bytesRead := byteutil.ReadAvailableBytesToBuffer(this.buffer, this.offset, data, offset, length)
	this.offset += bytesRead

	if this.offset == PROTOCOL_VERSION_PACKET_SIZE {
		version := binary.LittleEndian.Uint32(this.buffer)
		switch version {
			case 1:
				protocol.SetVersion(int(version))
				protocol.SetState(version1.PACKAGE_TYPE_STATE)

				break

			default:
				return bytesRead, errors.New("invalid protocol version " + strconv.Itoa(int(version)))
		}
	}

	return bytesRead, nil
}