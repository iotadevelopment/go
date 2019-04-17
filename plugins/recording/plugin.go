package recording

import (
    "github.com/iotadevelopment/go/packages/ixi"
    "github.com/iotadevelopment/go/plugins/gossip"
)

func configure() {
    neighborManager := gossip.GetNeighborManager()

    for _, neighbor := range neighborManager.GetStaticNeighbors() { setupNodeEventHandlers(neighbor) }
    for _, neighbor := range neighborManager.GetDynamicNeighbors() { setupNodeEventHandlers(neighbor) }

    neighborManager.Events.AddNeighbor.Attach(setupNodeEventHandlers)
    neighborManager.Events.RemoveNeighbor.Attach(tearNodeDownEventHandlers)
}

func run() {
    go writeData()
}

var PLUGIN = ixi.NewPlugin(configure, run)
