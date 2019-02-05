package gossip

import (
    "github.com/iotadevelopment/go/modules/gossip/protocol"
    "github.com/iotadevelopment/go/packages/ixi"
    "github.com/iotadevelopment/go/packages/network"
    "github.com/iotadevelopment/go/packages/transaction"
)

var gossipIXI = IXI()
var tcpServer = gossipIXI.GetTcpServer()

func configure() {
    tcpServer.Events.Connect.Attach(func(peer network.Peer) {
        gossipIXI.TriggerConnect(peer)

        gossipProtocol := protocol.NewProtocol().OnReceivePacketData(func(data []byte) {
            gossipIXI.TriggerReceivePacketData(peer, data)

            go parseTransaction(peer, data)
        })

        peer.OnReceiveData(func(data []byte) {
            gossipIXI.TriggerReceiveData(peer, data)

            gossipProtocol.ParseData(data)
        }).OnDisconnect(func() {
            gossipIXI.TriggerDisconnect(peer)
        }).OnError(func(err error) {
            gossipIXI.TriggerPeerError(peer, err)
        })
    })

    tcpServer.Events.Error.Attach(func(err error) {
        gossipIXI.TriggerError(err)
    })
}

func parseTransaction(peer network.Peer, data []byte) {
    transaction := transaction.FromBytes(data)

    gossipIXI.TriggerReceiveTransaction(peer, transaction)
}

func run() {
    go tcpServer.Listen(PORT_TCP.GetValue())
}

var PLUGIN = ixi.NewPlugin(configure, run)