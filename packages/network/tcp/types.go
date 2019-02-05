package tcp

import (
    "github.com/iotadevelopment/go/packages/network"
)

type serverEvents struct {
    Connect *peerConsumerEvent
    Error   *errorConsumerEvent
}

type peerConsumerEvent struct {
    callbacks map[uintptr]PeerConsumer
}

type errorConsumerEvent struct {
    callbacks map[uintptr]ErrorConsumer
}

type ErrorConsumer = func(e error)

type PeerConsumer = func(peer network.Peer)
