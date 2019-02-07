package tangle

import (
    "fmt"
    "github.com/iotadevelopment/go/packages/database"
    "github.com/iotadevelopment/go/packages/ixi"
    "github.com/iotadevelopment/go/packages/transaction"
    "github.com/iotadevelopment/go/plugins/gossip"
    "strconv"
)

var PLUGIN = ixi.NewPlugin(func() {
    transactionsDatabase := database.Get("transactions")

    counter := 0

    neighborManager := gossip.GetNeighborManager()
    neighborManager.Events.AddNeighbor.Attach(func(neighbor *gossip.Neighbor) {
        neighbor.Events.ReceiveTransaction.Attach(func(transaction *transaction.Transaction) {
            err := transactionsDatabase.Set([]byte(transaction.Hash.ToString() + strconv.Itoa(counter)), transaction.Bytes)
            if err != nil {
                fmt.Println(err)
            }
        })
    })
})
