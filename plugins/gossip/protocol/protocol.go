package protocol

import (
    "github.com/iotadevelopment/go/plugins/gossip/protocol/interfaces"
    "github.com/iotadevelopment/go/plugins/gossip/protocol/version1"
)

type protocolImplementation struct {
    availableVersions         map[int]map[int]interfaces.ProtocolState
    currentVersion            int
    currentState              interfaces.ProtocolState
    receivePacketDataHandlers []interfaces.DataConsumer
}

func NewProtocol() interfaces.Protocol {
    this := &protocolImplementation{
        availableVersions:         make(map[int]map[int]interfaces.ProtocolState),
        currentVersion:            0,
        currentState:              NewProtocolVersionState(),
        receivePacketDataHandlers: make([]interfaces.DataConsumer, 0),
    }

    version1.Register(this)

    return this
}

func (this *protocolImplementation) ParseData(data []byte) error {
    offset := 0
    length := len(data)
    for offset < length {
        readBytes, err := this.currentState.Consume(this, data, offset, length)
        offset += readBytes
        if err != nil {
            return err
        }
    }

    return nil
}

func (this *protocolImplementation) SetVersion(version int) {
    this.currentVersion = version
}

func (this *protocolImplementation) RegisterVersion(version int, stateMapping map[int]interfaces.ProtocolState) {
    this.availableVersions[version] = stateMapping
}

func (this *protocolImplementation) SetState(state int) {
    this.currentState = this.availableVersions[this.currentVersion][state]
}

func (this *protocolImplementation) OnReceivePacketData(callback interfaces.DataConsumer) interfaces.Protocol {
    this.receivePacketDataHandlers = append(this.receivePacketDataHandlers, callback)

    return this
}

func (this *protocolImplementation) TriggerReceivePacketData(data []byte) interfaces.Protocol {
    for _, callback := range this.receivePacketDataHandlers {
        callback(data)
    }

    return this
}