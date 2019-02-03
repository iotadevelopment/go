package gossip

func (this *IXIGossip) OnReceiveData(eventHandler DataConsumer) {
    this.receiveDataHandlers = append(this.receiveDataHandlers, eventHandler)
}

func (this *IXIGossip) OnReceiveTransactionData(eventHandler DataConsumer) {
    this.receiveTransactionDataHandlers = append(this.receiveDataHandlers, eventHandler)
}

func (this *IXIGossip) TriggerReceiveData(peer Peer, rawTransaction []byte) {
    for _, handler := range this.receiveDataHandlers {
        handler(peer, rawTransaction)
    }
}

func (this *IXIGossip) TriggerReceiveTransactionData(peer Peer, rawTransaction []byte) {
    for _, handler := range this.receiveTransactionDataHandlers {
        handler(peer, rawTransaction)
    }
}