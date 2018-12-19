package gossip

import (
	"ixi"
	"net"
)

type PeerManager struct {
	ixi         *ixi.IXI
	gossipPeers map[net.Conn]*Peer
}

func NewPeerManager(ixi *ixi.IXI) *PeerManager {
	this := &PeerManager{ixi, make(map[net.Conn]*Peer)}

	return this
}

func (this *PeerManager) AddPeer(conn net.Conn, peer *Peer)  {
	this.gossipPeers[conn] = peer
}

func (this *PeerManager) GetPeer(conn net.Conn) (*Peer, bool) {
	peer, exists := this.gossipPeers[conn]

	return peer, exists
}