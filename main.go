package main

import (
    "github.com/iotadevelopment/go/packages/ixi"
    "github.com/iotadevelopment/go/plugins/autopeering"
    "github.com/iotadevelopment/go/plugins/cli"
    "github.com/iotadevelopment/go/plugins/config"
    "github.com/iotadevelopment/go/plugins/gossip"
    "github.com/iotadevelopment/go/plugins/recording"
    "github.com/iotadevelopment/go/plugins/statusscreen"
    "github.com/iotadevelopment/go/plugins/tangle"
)

func main() {
    ixi.Load(
        // allow the node to be configured through a config file
        config.PLUGIN,

        // allow the node to be configure via cli parameters
        cli.PLUGIN,

        // add the gossip layer
        gossip.PLUGIN,

        // allow auto peering
        autopeering.PLUGIN,

        // add persistence layer
        tangle.PLUGIN,

        recording.PLUGIN,

        // show a status screen while the node is running
        statusscreen.PLUGIN,
    )

    ixi.Run()
}
