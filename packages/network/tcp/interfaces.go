package tcp

import (
    "github.com/iotadevelopment/go/packages/network"
)

type Server interface {
    Listen(port int) Server
    OnConnect(eventHandler PeerConsumer) Server
    OnError(eventHandler ErrorConsumer) Server
    TriggerConnect(peer network.Peer) Server
    TriggerError(err error) Server
}

type ErrorConsumer func(e error)

type PeerConsumer func(peer network.Peer)
