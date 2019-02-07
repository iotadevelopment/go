package gossip

import "github.com/iotadevelopment/go/packages/transaction"

type Callback = func()

type DataConsumer = func(data []byte)

type TransactionConsumer = func(transaction *transaction.Transaction)

type ErrorConsumer = func(err error)
