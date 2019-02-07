package gossip

import (
    "github.com/iotadevelopment/go/packages/ixi"
    "github.com/iotadevelopment/go/packages/network"
    "github.com/iotadevelopment/go/packages/network/tcp"
    "github.com/iotadevelopment/go/packages/transaction"
)

var tcpServer = tcp.NewServer()

func configure() {
    tcpServer.Events.Connect.Attach(func(conn network.Connection) {
        neighbor := NewNeighbour()

        neighbor.Events.IncomingConnection.Attach(func() {
            Events.Connect.Trigger(conn)
        })
        neighbor.Events.ReceiveData.Attach(func(data []byte) {
            Events.ReceiveData.Trigger(conn, data)
        })
        neighbor.Events.ReceiveTransactionData.Attach(func(data []byte) {
            Events.ReceiveTransactionData.Trigger(conn, data)
        })
        neighbor.Events.ReceiveTransaction.Attach(func(tx *transaction.Transaction) {
            Events.ReceiveTransaction.Trigger(conn, tx)
        })
        neighbor.Events.Disconnect.Attach(func() {
            Events.Disconnect.Trigger(conn)
        })
        neighbor.Events.Error.Attach(func(err error) {
            Events.PeerError.Trigger(conn, err)
        })

        neighbor.SetIncomingConnection(conn)
    })

    tcpServer.Events.Error.Attach(Events.Error.Trigger)
}

func run() {
    go tcpServer.Listen(*PORT_TCP.Value)
}

var PLUGIN = ixi.NewPlugin(configure, run)