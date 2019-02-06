package tcp

import (
    "github.com/iotadevelopment/go/packages/network"
    "net"
    "strconv"
)

type Server struct {
    Events serverEvents
}

func (this *Server) Listen(port int) *Server {
    l, err := net.Listen("tcp4", "0.0.0.0:"+strconv.Itoa(port))
    if err != nil {
        this.Events.Error.Trigger(err)

        return this
    }
    defer l.Close()

    for {
        c, err := l.Accept()
        if err != nil {
            this.Events.Error.Trigger(err)

            return this
        }

        peer := network.NewPeer("tcp", c)

        this.Events.Connect.Trigger(peer)

        go peer.HandleConnection()
    }

    return this
}

func NewServer() *Server {
    return &Server{
        Events: serverEvents{
            Connect: &peerConsumerEvent{make(map[uintptr]PeerConsumer)},
            Error: &errorConsumerEvent{make(map[uintptr]ErrorConsumer)},
        },
    }
}