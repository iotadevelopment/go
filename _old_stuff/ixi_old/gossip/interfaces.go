package gossip

type Peer interface {
}

type DataConsumer func(peer Peer, transactionData []byte)
