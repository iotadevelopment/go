package tcpprotocol

import (
    "github.com/iotadevelopment/go/packages/byteutils"
    "strconv"
)

//region portState /////////////////////////////////////////////////////////////////////////////////////////////////////

type portState struct {
    buffer []byte
    offset int
}

func (this *portState) Consume(protocol *protocol, data []byte, offset int, length int) (int, error) {
    bytesRead := byteutils.ReadAvailableBytesToBuffer(this.buffer, this.offset, data, offset, length)

    this.offset += bytesRead
    if this.offset == PORT_BYTES_COUNT {
        port, err := strconv.Atoi(string(this.buffer))
        if err != nil {
            protocol.Events.Error.Trigger(err)
        } else {
            protocol.Events.ReceivePortData.Trigger(port)
        }

        this.offset = 0

        protocol.currentState = protocol.transactionState
    }

    return bytesRead, nil
}

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
        transactionData := make([]byte, TRANSACTION_BYTES_COUNT)
        copy(transactionData, this.buffer)
        protocol.Events.ReceiveTransactionData.Trigger(transactionData)
        this.offset = 0

        protocol.currentState = protocol.requestState
    }

    return bytesRead, nil
}

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
        requestData := make([]byte, REQUEST_BYTES_COUNT)
        copy(requestData, this.buffer)
        protocol.Events.ReceiveTransactionRequestData.Trigger(requestData)
        this.offset = 0

        protocol.currentState = protocol.crc32State
    }

    return bytesRead, nil
}

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
        crc32Data := make([]byte, CRC32_BYTES_COUNT)
        copy(crc32Data, this.buffer)
        this.offset = 0

        protocol.currentState = protocol.transactionState
    }

    return bytesRead, nil
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////