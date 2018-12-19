package version1

import (
    "core/byteutil"
    "core/network/gossip/protocol/interfaces"
    "fmt"
    "ixi"
    "time"
)

type ParsedDataReceiver interface {
    TriggerReceivePackage([]byte)
}

var (
    receivedPackets = 0
    lastLogTime     = time.Now().UnixNano() / 1e6
)

type PackageState struct {
    ixi    *ixi.IXI
    buffer []byte
    offset int
}

func NewPackageState(ixi *ixi.IXI) *PackageState {
    this := &PackageState{ixi, make([]byte, PACKAGE_SIZE), 0}

    return this
}

func (this *PackageState) Reset() {
    this.offset = 0
}

func (this *PackageState) Consume(protocol interfaces.Protocol, data []byte, offset int, length int) (int, error) {
    bytesRead := byteutil.ReadAvailableBytesToBuffer(this.buffer, this.offset, data, offset, length)
    this.offset += bytesRead

    if this.offset == PACKAGE_SIZE {
        receivedPackets++

        this.ixi.Gossip.TriggerReceiveTransactionData(nil, this.buffer)

        //go model.NewTransactionFromBytes(&this.buffer)

        now := time.Now().UnixNano() / 1e6
        if now-lastLogTime >= 1000 {
            fmt.Println(receivedPackets, "tps")

            lastLogTime = now
            receivedPackets = 0
        }

        protocol.SetState(PACKAGE_TYPE_STATE)
        this.Reset()
    }

    return bytesRead, nil
}
