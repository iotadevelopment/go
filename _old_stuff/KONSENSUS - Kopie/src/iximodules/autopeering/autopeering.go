package autopeering

import (
    "core/network/gossip"
    "fmt"
    "ixi"
    "net"
)

type Peer interface {
    TriggerReceiveData(data *[]byte)
}

type AutoPeering struct {
    ixi         *ixi.IXI
    peerManager *gossip.PeerManager
}

func NewAutoPeering(ixi *ixi.IXI, manager *gossip.PeerManager) *AutoPeering {
    this := &AutoPeering{
        ixi:         ixi,
        peerManager: manager,
    }

    return this
}

func (this *AutoPeering) Start() {
    this.ixi.Network.OnClientConnect(func(c net.Conn) {
        peer := gossip.NewPeer(this.ixi)
        peer.SetConnection(c)

        this.peerManager.AddPeer(c, peer)

        fmt.Printf("Unknown client connected from %s\n", c.RemoteAddr().String())
        fmt.Println("Not enough gossip peers ... [ACCEPTING CONNECTION]")
    })

    this.ixi.Network.OnClientDisconnect(func(c net.Conn) {
        fmt.Printf("Client disconnected from %s\n", c.RemoteAddr().String())
    })

    this.ixi.Network.OnClientReceiveData(func(c net.Conn, data []byte) {
        if peer, exists := this.peerManager.GetPeer(c); exists {
            peer.TriggerReceiveData(data)
        }
    })

    this.ixi.Network.OnClientError(func(c net.Conn, err error) {
        fmt.Println("Error:", err)
    })
}
