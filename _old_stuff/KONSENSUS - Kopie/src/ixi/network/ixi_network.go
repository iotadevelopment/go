package network

type IXINetwork struct {
    errorHandlers             []ErrorConsumer
    clientConnectHandlers     []SocketConsumer
    clientDisconnectHandlers  []SocketConsumer
    clientReceiveDataHandlers []SocketDataConsumer
    clientErrorHandlers       []SocketErrorConsumer
}

func NewIXINetwork() *IXINetwork {
    this := &IXINetwork{}

    return this
}