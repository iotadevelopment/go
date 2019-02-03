package version1

import (
    "core/byteutil"
    "core/network/gossip/protocol/interfaces"
)

type PackageTypeState struct {}

func NewPackageTypeState() *PackageTypeState {
    this := &PackageTypeState{}

    return this
}

func (this *PackageTypeState) Consume(protocol interfaces.Protocol, data []byte, offset int, length int) (int, error) {
    readBytes := byteutil.ReadBytesIfAvailable(1, data, offset, length)
    if readBytes != nil {
        protocol.SetState(PACKAGE_STATE)

        return 1, nil
    }

    return 0, nil
}
