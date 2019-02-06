package gossip

import (
    "github.com/iotadevelopment/go/packages/network/tcp"
    "github.com/iotadevelopment/go/plugins/gossip/protocol"
    "github.com/iotadevelopment/go/packages/ixi"
    "github.com/iotadevelopment/go/packages/network"
    "github.com/iotadevelopment/go/packages/transaction"
)

var tcpServer = tcp.NewServer()

func configure() {
    tcpServer.Events.Connect.Attach(func(peer network.Peer) {
        Events.Connect.Trigger(peer)

        gossipProtocol := protocol.NewProtocol().OnReceivePacketData(func(data []byte) {
            Events.ReceivePacketData.Trigger(peer, data)

            go parseTransaction(peer, data)
        })

        peer.OnReceiveData(func(data []byte) {
            Events.ReceiveData.Trigger(peer, data)

            gossipProtocol.ParseData(data)
        }).OnDisconnect(func() {
            Events.Disconnect.Trigger(peer)
        }).OnError(func(err error) {
            Events.PeerError.Trigger(peer, err)
        })
    })

    tcpServer.Events.Error.Attach(Events.Error.Trigger)
}

func parseTransaction(peer network.Peer, data []byte) {
    Events.ReceiveTransaction.Trigger(peer, transaction.FromBytes(data))
}

func run() {
    go tcpServer.Listen(*PORT_TCP.Value)
}

var PLUGIN = ixi.NewPlugin(configure, run)