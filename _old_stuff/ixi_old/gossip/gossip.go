package gossip

type IXIGossip struct {
    receiveDataHandlers            []DataConsumer
    receiveTransactionDataHandlers []DataConsumer
}

func NewIXIGossip() *IXIGossip {
    this := &IXIGossip{}

    return this
}
