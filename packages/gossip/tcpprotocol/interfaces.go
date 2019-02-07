package tcpprotocol

type protocolState interface {
    Consume(protocol *protocol, data []byte, offset int, length int) (int, error)
}