package version1

import (
	"core/network/gossip/protocol/interfaces"
	"ixi"
)

const (
	// enum for the states
	PACKAGE_TYPE_STATE = 0
	PACKAGE_STATE = 1

	// other parameters
	PACKAGE_SIZE = 1604
)

func Register(ixi *ixi.IXI, protocol interfaces.Protocol) {
	stateMapping := make(map[int]interfaces.ProtocolState)

	stateMapping[PACKAGE_TYPE_STATE] = NewPackageTypeState()
	stateMapping[PACKAGE_STATE] = NewPackageState(ixi)

	protocol.RegisterVersion(1, stateMapping)
}