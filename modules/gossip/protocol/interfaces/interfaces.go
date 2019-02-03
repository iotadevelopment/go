package interfaces

type DataConsumer func(data []byte)

type ProtocolState interface {
    Consume(protocol Protocol, data []byte, offset int, length int) (int, error)
}

type Protocol interface {
    RegisterVersion(version int, stateMapping map[int]ProtocolState)
    SetState(state int)
    SetVersion(version int)
    ParseData(data []byte) error
    OnReceivePacketData(callback DataConsumer) Protocol
    TriggerReceivePacketData(data []byte) Protocol
}

