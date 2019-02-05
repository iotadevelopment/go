package gossip

import (
    "github.com/iotadevelopment/go/packages/transaction"
    "github.com/iotadevelopment/go/packages/network"
    "github.com/iotadevelopment/go/packages/network/tcp"
)

type PeerConsumer func(peer network.Peer)

type ErrorConsumer func(err error)

type PeerDataConsumer func(peer network.Peer, data []byte)

type PeerErrorConsumer func(peer network.Peer, err error)

type PeerTransactionConsumer func(peer network.Peer, transaction transaction.Transaction)

type GossipIXI interface {
    GetTcpServer() tcp.Server

    OnConnect(callback PeerConsumer) GossipIXI
    OnError(callback ErrorConsumer) GossipIXI
    OnReceiveData(callback PeerDataConsumer) GossipIXI
    OnDisconnect(callback PeerConsumer) GossipIXI
    OnPeerError(callback PeerErrorConsumer) GossipIXI
    OnReceivePacketData(callback PeerDataConsumer) GossipIXI
    OnReceiveTransaction(callback PeerTransactionConsumer) GossipIXI

    TriggerConnect(peer network.Peer) GossipIXI
    TriggerError(err error) GossipIXI
    TriggerReceiveData(peer network.Peer, data []byte) GossipIXI
    TriggerDisconnect(peer network.Peer) GossipIXI
    TriggerPeerError(peer network.Peer, err error) GossipIXI
    TriggerReceivePacketData(peer network.Peer, data []byte) GossipIXI
    TriggerReceiveTransaction(peer network.Peer, transaction transaction.Transaction) GossipIXI
}
