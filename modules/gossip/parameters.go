package gossip

import "github.com/iotadevelopment/go/packages/parameter"

var (
    PORT_TCP *parameter.IntParameter
    PORT_UDP *parameter.IntParameter
)

func initializeParameters(params *parameter.ParameterIXI) {
    PORT_TCP = params.AddInt("GOSSIP/PORT_TCP", 14625, "tcp port for incoming gossip messages")
    PORT_UDP = params.AddInt("GOSSIP/PORT_UDP", 14626, "udp port for incoming gossip messages")
}
