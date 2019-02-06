package gossip

var (
    Events = gossipEvents{
        Connect:            &peerEvent{make(map[uintptr]PeerConsumer)},
        Error:              &errorEvent{make(map[uintptr]ErrorConsumer)},
        ReceiveData:        &peerDataEvent{make(map[uintptr]PeerDataConsumer)},
        Disconnect:         &peerEvent{make(map[uintptr]PeerConsumer)},
        PeerError:          &peerErrorEvent{make(map[uintptr]PeerErrorConsumer)},
        ReceivePacketData:  &peerDataEvent{make(map[uintptr]PeerDataConsumer)},
        ReceiveTransaction: &peerTransactionEvent{make(map[uintptr]PeerTransactionConsumer)},
    }
)