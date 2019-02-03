package network

import "net"

type Peer interface {
    GetProtocol() string
    GetConnection() net.Conn
    OnReceiveData(callback DataConsumer) Peer
    OnDisconnect(callback Callback) Peer
    OnError(callback ErrorConsumer) Peer
    TriggerReceiveData(data []byte) Peer
    TriggerDisconnect() Peer
    TriggerError(err error) Peer
    HandleConnection()
}

type Callback func()

type ErrorConsumer func(err error)

type DataConsumer func(data []byte)
