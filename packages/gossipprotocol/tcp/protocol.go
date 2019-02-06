package tcp

type protocol struct {
    Events protocolEvents
    currentState protocolState
}

func New() *protocol {
    return &protocol{
        Events: protocolEvents{
            ReceivePortData:               &dataEvent{make(map[uintptr]DataConsumer)},
            ReceiveTransactionData:        &dataEvent{make(map[uintptr]DataConsumer)},
            ReceiveTransactionRequestData: &dataEvent{make(map[uintptr]DataConsumer)},
            Error:                         &errorEvent{make(map[uintptr]DataConsumer)},
        },
        currentState: port_state,
    }
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