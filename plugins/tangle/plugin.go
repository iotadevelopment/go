package tangle

import (
    "fmt"
    "github.com/iotadevelopment/go/packages/database"
    "github.com/iotadevelopment/go/packages/ixi"
    "github.com/iotadevelopment/go/packages/transaction"
    "github.com/iotadevelopment/go/plugins/gossip"
)

var transactionsDatabase, err = database.Get("transactions")

func StoreTransaction(transaction *transaction.Transaction) {
    transactionHash := []byte(transaction.Hash.ToString())

    _, err := transactionsDatabase.Get(transactionHash)
    if err == database.ErrKeyNotFound {
        err := transactionsDatabase.Set(transactionHash, transaction.Bytes)
        if err != nil {
            fmt.Println(err)
        }
    }
}

func setupEventHandlers(neighbor *gossip.Neighbor) {
    neighbor.Events.ReceiveTransaction.Attach(StoreTransaction)
}

func tearDownEventHandlers(neighbor *gossip.Neighbor) {
    neighbor.Events.ReceiveTransaction.Detach(StoreTransaction)
}

func configure() {
    neighborManager := gossip.GetNeighborManager()

    // setup event handlers for existing neighbors
    for _, neighbor := range neighborManager.GetStaticNeighbors() {
        setupEventHandlers(neighbor)
    }
    for _, neighbor := range neighborManager.GetDynamicNeighbors() {
        setupEventHandlers(neighbor)
    }

    // setup event handlers for new neighbors
    neighborManager.Events.AddNeighbor.Attach(setupEventHandlers)
    neighborManager.Events.RemoveNeighbor.Attach(tearDownEventHandlers)
}

var PLUGIN = ixi.NewPlugin(configure)
