package tcpprotocol

type protocolEvents struct {
    ReceiveTransactionData        *dataEvent
    ReceiveTransactionRequestData *dataEvent
    Error                         *errorEvent
}
