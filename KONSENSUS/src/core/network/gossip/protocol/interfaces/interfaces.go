package interfaces

type ProtocolState interface {
	Consume(protocol Protocol, data []byte, offset int, length int) (int, error)
}

type Protocol interface {
	RegisterVersion(version int, stateMapping map[int]ProtocolState)
	SetState(state int)
	SetVersion(version int)
	ParseData(data []byte, length int) error
}
