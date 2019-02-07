package gossip

import (
    "fmt"
    "github.com/iotadevelopment/go/packages/network"
    "github.com/iotadevelopment/go/packages/transaction"
    "github.com/iotadevelopment/go/plugins/gossip/tcpprotocol"
    "net"
    "strconv"
)

type Neighbor struct {
    Events   neighborEvents
    protocol string
    ip       string
    port     int
    in       network.Connection
    out      network.Connection
}

func newNeighbour() *Neighbor {
    neighbor := &Neighbor{
        Events: neighborEvents{
            IncomingConnection:     &callbackEvent{make(map[uintptr]Callback)},
            ReceiveData:            &dataEvent{make(map[uintptr]DataConsumer)},
            ReceiveTransactionData: &dataEvent{make(map[uintptr]DataConsumer)},
            ReceiveTransaction:     &transactionEvent{make(map[uintptr]TransactionConsumer)},
            Disconnect:             &callbackEvent{make(map[uintptr]Callback)},
            Error:                  &errorEvent{make(map[uintptr]ErrorConsumer)},
        },
    }

    return neighbor
}

func (this *Neighbor) Connect() {
    conn, err := net.Dial("tcp", this.ip + ":" + strconv.Itoa(this.port))
    if err != nil {
        this.Events.Error.Trigger(err)
        return
    }

    connection := network.NewPeer("tcp", conn)

    this.SetOutgoingConnection(connection)

    go connection.HandleConnection()
}

func (this *Neighbor) SetOutgoingConnection(conn network.Connection) {
    // store connection specific information
    this.out = conn

    // write the port (according to tcp protocol)
    conn.Write([]byte(fmt.Sprintf("%010d", *PORT_TCP.Value)))

    // dispatch raw low level events
    conn.OnReceiveData(this.Events.ReceiveData.Trigger)
    conn.OnDisconnect(this.Events.Disconnect.Trigger)
    conn.OnError(this.Events.Error.Trigger)
}

func (this *Neighbor) SetIncomingConnection(conn network.Connection) {
    switch conn.GetProtocol() {
    case "tcp":
        // store connection specific information
        this.in = conn
        if addr, ok := conn.GetConnection().RemoteAddr().(*net.TCPAddr); ok {
            this.ip = addr.IP.String()
        }

        // trigger connect event
        this.Events.IncomingConnection.Trigger()

        // dispatch raw low level events
        conn.OnReceiveData(this.Events.ReceiveData.Trigger)
        conn.OnDisconnect(this.Events.Disconnect.Trigger)
        conn.OnError(this.Events.Error.Trigger)

        // dispatch high level events
        tcpProtocol := tcpprotocol.New()
        tcpProtocol.Events.ReceivePortData.Attach(func(port int) {
            this.port = port

            if this.out == nil {
                this.Connect()
            }
        })
        tcpProtocol.Events.ReceiveTransactionData.Attach(func(data []byte) {
            go func() {
                this.Events.ReceiveTransaction.Trigger(transaction.FromBytes(data))
            }()
        })
        tcpProtocol.Events.ReceiveTransactionData.Attach(this.Events.ReceiveTransactionData.Trigger)
        tcpProtocol.Events.Error.Attach(this.Events.Error.Trigger)

        // start the processing of the protocol
        conn.OnReceiveData(tcpProtocol.ParseData)

        break;
    }
}