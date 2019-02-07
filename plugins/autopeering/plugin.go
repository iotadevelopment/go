package autopeering

import (
    "github.com/iotadevelopment/go/packages/ixi"
    "github.com/iotadevelopment/go/plugins/gossip"
)

var PLUGIN = ixi.NewPlugin(func() {
    neighbourManager := gossip.GetNeighborManager()

    gossip.Events.ConnectUnknownNeighbor.Attach(func(neighbor *gossip.Neighbor) {
        neighbourManager.AddNeighbor(neighbor, false)
    })
}, func() {

})
