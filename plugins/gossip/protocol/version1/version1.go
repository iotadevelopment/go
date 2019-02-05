package version1

import "github.com/iotadevelopment/go/plugins/gossip/protocol/interfaces"

const (
    // enum for the states
    PACKAGE_TYPE_STATE = 0
    PACKAGE_STATE = 1

    // other parameters
    PACKAGE_SIZE = 1604
)

func Register(protocol interfaces.Protocol) {
    stateMapping := make(map[int]interfaces.ProtocolState)

    stateMapping[PACKAGE_TYPE_STATE] = NewPackageTypeState()
    stateMapping[PACKAGE_STATE] = NewPackageState()

    protocol.RegisterVersion(1, stateMapping)
}
