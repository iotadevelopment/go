package tangle

import (
    "fmt"
    "github.com/iotadevelopment/go/plugins/gossip"
    "github.com/iotadevelopment/go/packages/database"
    "github.com/iotadevelopment/go/packages/ixi"
    "github.com/iotadevelopment/go/packages/transaction"
    "github.com/iotadevelopment/go/packages/network"
    "strconv"
)

var PLUGIN = ixi.NewPlugin(func() {
    transactionsDatabase := database.Get("transactions")

    counter := 0

    gossip.Events.ReceiveTransaction.Attach(func(peer network.Peer, transaction *transaction.Transaction) {
        counter++

        err := transactionsDatabase.Set([]byte(transaction.Hash.ToString() + strconv.Itoa(counter)), transaction.Bytes)
        if err != nil {
            fmt.Println(err)
        }
    })
})
