package gossip

import (
    "errors"
    "github.com/iotadevelopment/go/packages/ixi"
    "github.com/iotadevelopment/go/packages/network"
    "github.com/iotadevelopment/go/packages/network/tcp"
    "github.com/iotadevelopment/go/plugins/gossip/tcpprotocol"
    "net"
    "strconv"
)

var tcpServer = tcp.NewServer()

var neighborManager = &NeighborManager{
    Events: neighborManagerEvents{
        AddNeighbor:    &neighborEvent{map[uintptr]NeighborConsumer{}},
        RemoveNeighbor: &neighborEvent{map[uintptr]NeighborConsumer{}},
    },

    staticNeighbors:  make(map[string]*Neighbor, 0),
    dynamicNeighbors: make(map[string]*Neighbor, 0),
}

func readPort(conn network.Connection) (int, error) {
    portBytes := make([]byte, tcpprotocol.PORT_BYTES_COUNT)

    readOffset := 0
    for readOffset < tcpprotocol.PORT_BYTES_COUNT {
        receiveBuffer := make([]byte, tcpprotocol.PORT_BYTES_COUNT - readOffset)
        if bytesRead, err := conn.GetConnection().Read(receiveBuffer); err != nil {
            return 0, err
        } else {
            copy(portBytes[readOffset:], receiveBuffer[:bytesRead])

            readOffset += bytesRead
        }
    }

    if port, err := strconv.Atoi(string(portBytes)); err != nil {
        return 0, err
    } else {
        if port < 1 || port > 65535 {
            return 0, errors.New("invalid remote port (needs to be between 1 and 65535)")
        }

        return port, nil
    }
}

func configure() {
    tcpServer.Events.Connect.Attach(func(conn network.Connection) {
        host := conn.GetConnection().RemoteAddr().(*net.TCPAddr).IP.String()

        if  port, err := readPort(conn); err != nil {
            Events.Error.Trigger(err)
        } else {
            address := host + ":" + strconv.Itoa(port)

            neighbor := neighborManager.GetNeighbor(address)
            if neighbor == nil {
                neighbor = newNeighbour(host, port)

                Events.ConnectUnknownNeighbor.Trigger(neighbor)

                neighbor.SetIncomingConnection(conn)
            } else {
                neighbor.SetIncomingConnection(conn)
            }
        }
    })

    neighborManager.Events.AddNeighbor.Attach(func(neighbor *Neighbor) {
        neighbor.Events.IncomingConnection.Attach(func() {
            go neighbor.ProcessIncomingConnection()
        })
    })

    tcpServer.Events.Error.Attach(Events.Error.Trigger)
}

func run() {
    // add neighbors to the neighbor manager
    // ...

    go tcpServer.Listen(*PORT_TCP.Value)
    go neighborManager.LaunchConnections()
}

var PLUGIN = ixi.NewPlugin(configure, run)