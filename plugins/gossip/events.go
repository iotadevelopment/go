package gossip

type gossipEvents struct {
    Connect                *peerEvent
    Error                  *errorEvent
    ReceiveData            *peerDataEvent
    Disconnect             *peerEvent
    PeerError              *peerErrorEvent
    ReceiveTransactionData *peerDataEvent
    ReceiveTransaction     *peerTransactionEvent
}

type neighborEvents struct {
    IncomingConnection     *callbackEvent
    ReceiveData            *dataEvent
    ReceiveTransactionData *dataEvent
    ReceiveTransaction     *transactionEvent
    Disconnect             *callbackEvent
    Error                  *errorEvent
}
