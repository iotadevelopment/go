package ixi

import (
    "github.com/iotadevelopment/go/_old_stuff/ixi_olduff/ixi_old/gossip"
    "github.com/iotadevelopment/go/_old_stuff/ixi_olduff/ixi_old/network"
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