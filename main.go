package main

import (
	"github.com/iotadevelopment/go/modules/cli"
	"github.com/iotadevelopment/go/modules/config"
	"github.com/iotadevelopment/go/modules/gossip"
	"github.com/iotadevelopment/go/modules/statusscreen"
	"github.com/iotadevelopment/go/modules/tangle"
	"github.com/iotadevelopment/go/packages/ixi"
)

func main() {
	ixi.Load(
		// allow the node to be configured through a config file
		config.PLUGIN,

		// allow the node to be configure via cli parameters
		cli.PLUGIN,

		// add the gossip layer
		gossip.PLUGIN,

		// add persistence layer
		tangle.PLUGIN,

		// show a status screen while the node is running
		statusscreen.PLUGIN,
	)

	ixi.Run()
}
