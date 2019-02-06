package gossip

import (
    "fmt"
    tcpGossip "github.com/iotadevelopment/go/packages/gossipprotocol/tcp"
    "github.com/iotadevelopment/go/packages/ixi"
    "github.com/iotadevelopment/go/packages/network"
    "github.com/iotadevelopment/go/packages/network/tcp"
    "github.com/iotadevelopment/go/packages/transaction"
    "net"
    "strconv"
)

var tcpServer = tcp.NewServer()

func configure() {
    tcpServer.Events.Connect.Attach(func(peer network.Peer) {
        Events.Connect.Trigger(peer)

        gossipProtocol := tcpGossip.New()

        gossipProtocol.Events.ReceivePortData.Attach(func(data []byte) {
            port, _ := strconv.Atoi(string(data))

            conn, err := net.Dial("tcp", "95.216.33.102:" + strconv.Itoa(port))
            if err != nil {
                panic(err)
            }

            _, err = conn.Write([]byte(fmt.Sprintf("%010d", *PORT_TCP.Value)))
            if err != nil {
                panic(err)
            }
        })

        gossipProtocol.Events.ReceiveTransactionData.Attach(func(data []byte) {
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