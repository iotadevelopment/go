package gossip

import (
    "github.com/iotadevelopment/go/packages/transaction"
    "github.com/iotadevelopment/go/packages/network"
    "github.com/iotadevelopment/go/packages/network/tcp"
)

type gossipIXIImplementation struct {
    tcpServer                  *tcp.Server
    connectHandlers            []PeerConsumer
    errorHandlers              []ErrorConsumer
    receiveDataHandlers        []PeerDataConsumer
    disconnectHandlers         []PeerConsumer
    peerErrorHandlers          []PeerErrorConsumer
    receivePacketDataHandlers  []PeerDataConsumer
    receiveTransactionHandlers []PeerTransactionConsumer
}

var globalInstance *gossipIXIImplementation = nil

func IXI() GossipIXI {
    if globalInstance == nil {
        globalInstance = &gossipIXIImplementation{
            tcpServer: tcp.NewServer(),
        }
    }

    return globalInstance
}

func (this *gossipIXIImplementation) GetTcpServer() *tcp.Server {
    return this.tcpServer
}

func (this *gossipIXIImplementation) OnConnect(callback PeerConsumer) GossipIXI {
    this.connectHandlers = append(this.connectHandlers, callback)

    return this
}

func (this *gossipIXIImplementation) OnError(callback ErrorConsumer) GossipIXI {
    this.errorHandlers = append(this.errorHandlers, callback)

    return this
}

func (this *gossipIXIImplementation) OnReceiveData(eventHandler PeerDataConsumer) GossipIXI {
    this.receiveDataHandlers = append(this.receiveDataHandlers, eventHandler)

    return this
}

func (this *gossipIXIImplementation) OnDisconnect(callback PeerConsumer) GossipIXI {
    this.disconnectHandlers = append(this.disconnectHandlers, callback)

    return this
}

func (this *gossipIXIImplementation) OnPeerError(callback PeerErrorConsumer) GossipIXI {
    this.peerErrorHandlers = append(this.peerErrorHandlers, callback)

    return this
}

func (this *gossipIXIImplementation) OnReceivePacketData(eventHandler PeerDataConsumer) GossipIXI {
    this.receivePacketDataHandlers = append(this.receivePacketDataHandlers, eventHandler)

    return this
}

func (this *gossipIXIImplementation) OnReceiveTransaction(callback PeerTransactionConsumer) GossipIXI {
    this.receiveTransactionHandlers = append(this.receiveTransactionHandlers, callback)

    return this
}

func (this *gossipIXIImplementation) TriggerConnect(peer network.Peer) GossipIXI {
    for _, handler := range this.connectHandlers {
        handler(peer)
    }

    return this
}

func (this *gossipIXIImplementation) TriggerError(err error) GossipIXI {
    for _, handler := range this.errorHandlers {
        handler(err)
    }

    return this
}

func (this *gossipIXIImplementation) TriggerReceiveData(peer network.Peer, data []byte) GossipIXI {
    for _, handler := range this.receiveDataHandlers {
        handler(peer, data)
    }

    return this
}

func (this *gossipIXIImplementation) TriggerDisconnect(peer network.Peer) GossipIXI {
    for _, handler := range this.disconnectHandlers {
        handler(peer)
    }

    return this
}

func (this *gossipIXIImplementation) TriggerPeerError(peer network.Peer, err error) GossipIXI {
    for _, handler := range this.peerErrorHandlers {
        handler(peer, err)
    }

    return this
}

func (this *gossipIXIImplementation) TriggerReceivePacketData(peer network.Peer, data []byte) GossipIXI {
    for _, handler := range this.receivePacketDataHandlers {
        handler(peer, data)
    }

    return this
}

func (this *gossipIXIImplementation) TriggerReceiveTransaction(peer network.Peer, transaction transaction.Transaction) GossipIXI {
    for _, handler := range this.receiveTransactionHandlers {
        handler(peer, transaction)
    }

    return this
}
