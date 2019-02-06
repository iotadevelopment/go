package gossip

import (
    "github.com/iotadevelopment/go/packages/network"
    "github.com/iotadevelopment/go/packages/transaction"
)

type PeerConsumer = func(peer network.Peer)

type ErrorConsumer = func(err error)

type PeerDataConsumer = func(peer network.Peer, data []byte)

type PeerErrorConsumer = func(peer network.Peer, err error)

type PeerTransactionConsumer = func(peer network.Peer, transaction *transaction.Transaction)