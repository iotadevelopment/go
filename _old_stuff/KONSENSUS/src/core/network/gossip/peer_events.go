package gossip

func (this *Peer) OnReceiveData(eventHandler PeerDataConsumer) *Peer {
	this.receiveDataHandlers = append(this.receiveDataHandlers, eventHandler)

	return this
}

func (this *Peer) TriggerReceiveData(data []byte) {
	for _, handler := range this.receiveDataHandlers {
		handler(data)
	}
}