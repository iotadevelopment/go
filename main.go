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
		config.MODULE,

		// allow the node to be configure via cli parameters
		cli.MODULE,

		// add the gossip layer
		gossip.MODULE,

		tangle.MODULE,

		// show a banner when the node has started
		statusscreen.MODULE,
	)

	ixi.Run()
}
