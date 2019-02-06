package tcp

import "github.com/iotadevelopment/go/packages/byteutils"

//region portState /////////////////////////////////////////////////////////////////////////////////////////////////////

type portState struct {
    buffer []byte
    offset int
}

func (this *portState) Consume(protocol *protocol, data []byte, offset int, length int) (int, error) {
    bytesRead := byteutils.ReadAvailableBytesToBuffer(this.buffer, this.offset, data, offset, length)

    this.offset += bytesRead
    if this.offset == PORT_BYTES_COUNT {
        portData := make([]byte, PORT_BYTES_COUNT)
        copy(portData, this.buffer)

        protocol.Events.ReceivePortData.Trigger(portData)
        protocol.currentState = transaction_state
    }

    return bytesRead, nil
}

var port_state = &portState{make([]byte, PORT_BYTES_COUNT), 0}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region transactionState //////////////////////////////////////////////////////////////////////////////////////////////

type transactionState struct {
    buffer []byte
    offset int
}

func (this *transactionState) Consume(protocol *protocol, data []byte, offset int, length int) (int, error) {
    bytesRead := byteutils.ReadAvailableBytesToBuffer(this.buffer, this.offset, data, offset, length)

    this.offset += bytesRead
    if this.offset == TRANSACTION_BYTES_COUNT {
        protocol.currentState = request_state
    }

    return bytesRead, nil
}

var transaction_state = &transactionState{make([]byte, TRANSACTION_BYTES_COUNT), 0}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region requestState //////////////////////////////////////////////////////////////////////////////////////////////////

type requestState struct {
    buffer []byte
    offset int
}

func (this *requestState) Consume(protocol *protocol, data []byte, offset int, length int) (int, error) {
    bytesRead := byteutils.ReadAvailableBytesToBuffer(this.buffer, this.offset, data, offset, length)

    this.offset += bytesRead
    if this.offset == REQUEST_BYTES_COUNT {
        protocol.currentState = crc32_state
    }

    return bytesRead, nil
}

var request_state = &requestState{make([]byte, REQUEST_BYTES_COUNT), 0}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region crc32State ////////////////////////////////////////////////////////////////////////////////////////////////////

type crc32State struct {
    buffer []byte
    offset int
}

func (this *crc32State) Consume(protocol *protocol, data []byte, offset int, length int) (int, error) {
    bytesRead := byteutils.ReadAvailableBytesToBuffer(this.buffer, this.offset, data, offset, length)

    this.offset += bytesRead
    if this.offset == CRC32_BYTES_COUNT {
        transactionData := make([]byte, TRANSACTION_BYTES_COUNT)
        copy(transactionData, transaction_state.buffer)
        transaction_state.offset = 0

        requestData := make([]byte, REQUEST_BYTES_COUNT)
        copy(requestData, request_state.buffer)
        request_state.offset = 0

        crc32Data := make([]byte, CRC32_BYTES_COUNT)
        copy(crc32Data, this.buffer)
        this.offset = 0

        protocol.Events.ReceiveTransactionData.Trigger(transactionData)
        protocol.Events.ReceiveTransactionRequestData.Trigger(requestData)

        protocol.currentState = transaction_state
    }

    return bytesRead, nil
}

var crc32_state = &requestState{make([]byte, CRC32_BYTES_COUNT), 0}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////