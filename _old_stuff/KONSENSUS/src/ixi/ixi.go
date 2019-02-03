package ixi

import (
    "ixi/gossip"
    "ixi/network"
)

type IXI struct {
    Network *network.IXINetwork
    Gossip  *gossip.IXIGossip
}

func NewIXI() *IXI {
    this := &IXI{
        Network: network.NewIXINetwork(),
        Gossip:  gossip.NewIXIGossip(),
    }

    return this
}
