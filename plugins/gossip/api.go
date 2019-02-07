package gossip

import "github.com/iotadevelopment/go/packages/network/tcp"

var (
    Events = gossipEvents{
        ConnectUnknownNeighbor: &neighborEvent{make(map[uintptr]NeighborConsumer)},
        Error:                  &errorEvent{make(map[uintptr]ErrorConsumer)},
    }
)

func GetTCPServer() *tcp.Server {
    return tcpServer
}

func GetNeighborManager() *NeighborManager {
    return neighborManager
}