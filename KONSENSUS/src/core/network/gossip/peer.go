package gossip

import (
    "core/network/gossip/protocol"
    "ixi"
    "net"
)

type SocketDataConsumer func(data []byte)

type Peer struct {
    ixi                    *ixi.IXI
    protocol               *protocol.Protocol
    conn                   net.Conn
    receiveDataHandlers    []SocketDataConsumer
    receivePacketsHandlers []SocketDataConsumer
}

func NewPeer(ixi *ixi.IXI) *Peer {
    this := &Peer{}

    this.ixi = ixi
    this.protocol = protocol.NewProtocol(ixi, this)

    this.OnReceiveData(func(data []byte) {
        this.ixi.Gossip.TriggerReceiveData(this, nil)

        this.protocol.ParseData(data, len(data))
    })

    return this
}

func (this *Peer) SetConnection(conn net.Conn) {
    this.conn = conn
}

func (this *Peer) Connect(address string) {

}
