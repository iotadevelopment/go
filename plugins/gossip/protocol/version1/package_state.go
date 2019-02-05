package version1

import (
    "github.com/iotadevelopment/go/plugins/gossip/protocol/interfaces"
    "github.com/iotadevelopment/go/packages/byteutils"
)

type ParsedDataReceiver interface {
    TriggerReceivePackage([]byte)
}

type PackageState struct {
    buffer []byte
    offset int
}

func NewPackageState() *PackageState {
    this := &PackageState{make([]byte, PACKAGE_SIZE), 0}

    return this
}

func (this *PackageState) Reset() {
    this.offset = 0
}

func (this *PackageState) Consume(protocol interfaces.Protocol, data []byte, offset int, length int) (int, error) {
    bytesRead := byteutils.ReadAvailableBytesToBuffer(this.buffer, this.offset, data, offset, length)
    this.offset += bytesRead

    if this.offset == PACKAGE_SIZE {
        transactionData := make([]byte, len(this.buffer))
        copy(transactionData, this.buffer)

        protocol.TriggerReceivePacketData(transactionData)
        protocol.SetState(PACKAGE_TYPE_STATE)
        this.Reset()
    }

    return bytesRead, nil
}

