package tcp

import (
    "github.com/iotadevelopment/go/packages/network"
    "net"
    "strconv"
)

type serverImplementation struct {
    clientConnectHandlers []PeerConsumer
    errorHandlers         []ErrorConsumer
}

func NewServer() Server {
    return &serverImplementation{}
}

func (this *serverImplementation) Listen(port int) Server {
    l, err := net.Listen("tcp4", "127.0.0.1:"+strconv.Itoa(port))
    if err != nil {
        this.TriggerError(err)

        return this
    }
    defer l.Close()

    for {
        c, err := l.Accept()
        if err != nil {
            this.TriggerError(err)

            return this
        }

        peer := network.NewPeer("tcp", c)

        this.TriggerConnect(peer)

        go peer.HandleConnection()
    }

    return this
}

func (this *serverImplementation) OnConnect(eventHandler PeerConsumer) Server {
    this.clientConnectHandlers = append(this.clientConnectHandlers, eventHandler)

    return this
}

func (this *serverImplementation) OnError(eventHandler ErrorConsumer) Server {
    this.errorHandlers = append(this.errorHandlers, eventHandler)

    return this
}

func (this *serverImplementation) TriggerConnect(peer network.Peer) Server {
    for _, onConnectHandler := range this.clientConnectHandlers {
        onConnectHandler(peer)
    }

    return this
}

func (this *serverImplementation) TriggerError(err error) Server {
    for _, handler := range this.errorHandlers {
        handler(err)
    }

    return this
}
