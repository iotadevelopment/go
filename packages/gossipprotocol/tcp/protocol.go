package tcp

type protocol struct {
    Events           protocolEvents
    currentState     protocolState
    portState        *portState
    transactionState *transactionState
    requestState     *requestState
    crc32State       *crc32State
}

func New() *protocol {
    protocol := &protocol{
        Events: protocolEvents{
            ReceivePortData:               &dataEvent{make(map[uintptr]DataConsumer)},
            ReceiveTransactionData:        &dataEvent{make(map[uintptr]DataConsumer)},
            ReceiveTransactionRequestData: &dataEvent{make(map[uintptr]DataConsumer)},
            Error:                         &errorEvent{make(map[uintptr]DataConsumer)},
        },

        portState:        &portState{make([]byte, PORT_BYTES_COUNT), 0},
        transactionState: &transactionState{make([]byte, TRANSACTION_BYTES_COUNT), 0},
        requestState:     &requestState{make([]byte, REQUEST_BYTES_COUNT), 0},
        crc32State:       &crc32State{make([]byte, CRC32_BYTES_COUNT), 0},
    }

    protocol.currentState = protocol.portState

    return protocol
}

func (this *protocol) ParseData(data []byte) error {
    offset := 0
    length := len(data)
    for offset < length {
        readBytes, err := this.currentState.Consume(this, data, offset, length)
        offset += readBytes
        if err != nil {
            return err
        }
    }

    return nil
}