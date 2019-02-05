package gossip

import "github.com/iotadevelopment/go/modules/parameter"

var (
    PORT_TCP = parameter.AddInt("GOSSIP/PORT_TCP", 14625, "tcp port for incoming gossip messages")
    PORT_UDP = parameter.AddInt("GOSSIP/PORT_UDP", 14626, "udp port for incoming gossip messages")
)
