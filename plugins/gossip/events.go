package gossip

type gossipEvents struct {
    ConnectUnknownNeighbor *neighborEvent
    Error                  *errorEvent
}

type neighborManagerEvents struct {
    AddNeighbor    *neighborEvent
    RemoveNeighbor *neighborEvent
}

type neighborEvents struct {
    IncomingConnection     *callbackEvent
    ReceiveData            *dataEvent
    ReceiveTransactionData *dataEvent
    ReceiveTransaction     *transactionEvent
    Disconnect             *callbackEvent
    Error                  *errorEvent
}
