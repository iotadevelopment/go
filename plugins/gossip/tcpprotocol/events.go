package tcpprotocol

type protocolEvents struct {
    ReceivePortData               *intEvent
    ReceiveTransactionData        *dataEvent
    ReceiveTransactionRequestData *dataEvent
    Error                         *errorEvent
}
