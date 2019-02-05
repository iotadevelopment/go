package tcp

import (
    "github.com/iotadevelopment/go/packages/network"
)

type ErrorConsumer = func(e error)

type PeerConsumer = func(peer network.Peer)
