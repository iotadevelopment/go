package protocol

import (
    "core/network/gossip/protocol/interfaces"
    "core/network/gossip/protocol/version1"
    "ixi"
)

const (
    VERSION_TYPE_STATE = 0
)

type Peer interface {

}

type Protocol struct {
    peer              Peer
    availableVersions map[int]map[int]interfaces.ProtocolState
    currentVersion    int
    currentState      interfaces.ProtocolState
}

func NewProtocol(ixi *ixi.IXI, peer Peer) *Protocol {
    this := &Protocol{
        peer,
        make(map[int]map[int]interfaces.ProtocolState),
        0,
        NewProtocolVersionState(),
    }

    version1.Register(ixi, this)

    return this
}

func (this *Protocol) Start(dataChannel <-chan []byte) {
    for receivedData := range dataChannel {
        this.ParseData(receivedData, len(receivedData))
    }
}

func (this *Protocol) ParseData(data []byte, length int) error {
    offset := 0
    for offset < length {
        readBytes, err := this.currentState.Consume(this, data, offset, length)
        offset += readBytes
        if err != nil {
            return err
        }
    }

    return nil
}

func (this *Protocol) SetVersion(version int) {
    this.currentVersion = version
}

func (this *Protocol) RegisterVersion(version int, stateMapping map[int]interfaces.ProtocolState) {
    this.availableVersions[version] = stateMapping
}

func (this *Protocol) SetState(state int) {
    this.currentState = this.availableVersions[this.currentVersion][state]
}
