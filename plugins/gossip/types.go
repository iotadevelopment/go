package gossip

import (
    "github.com/iotadevelopment/go/packages/network"
    "github.com/iotadevelopment/go/packages/transaction"
)

type PeerConsumer = func(peer network.Connection)

type ErrorConsumer = func(err error)

type PeerDataConsumer = func(peer network.Connection, data []byte)

type PeerErrorConsumer = func(peer network.Connection, err error)

type PeerTransactionConsumer = func(peer network.Connection, transaction *transaction.Transaction)