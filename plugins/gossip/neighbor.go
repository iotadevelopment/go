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
    host     string
    port     int
    in       network.Connection
    out      network.Connection
}

func newNeighbour(host string, port int) *Neighbor {
    neighbor := &Neighbor{
        Events: neighborEvents{
            IncomingConnection:     &callbackEvent{make(map[uintptr]Callback)},
            ReceiveData:            &dataEvent{make(map[uintptr]DataConsumer)},
            ReceiveTransactionData: &dataEvent{make(map[uintptr]DataConsumer)},
            ReceiveTransaction:     &transactionEvent{make(map[uintptr]TransactionConsumer)},
            Disconnect:             &callbackEvent{make(map[uintptr]Callback)},
            Error:                  &errorEvent{make(map[uintptr]ErrorConsumer)},
        },

        host: host,
        port: port,
    }

    return neighbor
}

func (this *Neighbor) GetAddress() string {
    return this.host + ":" + strconv.Itoa(this.port)
}

func (this *Neighbor) ProcessIncomingConnection() {
    // dispatch raw low level events
    this.in.OnReceiveData(this.Events.ReceiveData.Trigger)
    this.in.OnDisconnect(this.Events.Disconnect.Trigger)
    this.in.OnError(this.Events.Error.Trigger)

    // create tcp protocol for high level events
    tcpProtocol := tcpprotocol.New()
    tcpProtocol.Events.ReceiveTransactionData.Attach(func(data []byte) {
        go func() {
            this.Events.ReceiveTransaction.Trigger(transaction.FromBytes(data))
        }()
    })
    tcpProtocol.Events.ReceiveTransactionData.Attach(this.Events.ReceiveTransactionData.Trigger)
    tcpProtocol.Events.Error.Attach(this.Events.Error.Trigger)

    // launch tcp protocol parsing
    this.in.OnReceiveData(tcpProtocol.ParseData)

    // start reading from the connection
    this.in.HandleConnection()
}

func (this *Neighbor) ProcessOutgoingConnection() {
    // write the port (according to tcp protocol)
    this.out.Write([]byte(fmt.Sprintf("%010d", *PORT_TCP.Value)))

    // dispatch raw low level events
    this.out.OnReceiveData(this.Events.ReceiveData.Trigger)
    this.out.OnDisconnect(func() {
        this.Events.Disconnect.Trigger()

        this.out = nil
    })
    this.out.OnError(this.Events.Error.Trigger)

    this.out.HandleConnection()
}

func (this *Neighbor) Connect() {
    if this.out == nil {
        conn, err := net.Dial("tcp", this.GetAddress())
        if err != nil {
            this.Events.Error.Trigger(err)

            return
        }

        this.SetOutgoingConnection(network.NewPeer("tcp", conn))

        go this.ProcessOutgoingConnection()
    }
}

func (this *Neighbor) SetOutgoingConnection(conn network.Connection) {
    // store connection specific information
    this.out = conn
}

func (this *Neighbor) SetIncomingConnection(conn network.Connection) {
    // store connection specific information
    this.in = conn

    // trigger connect event
    this.Events.IncomingConnection.Trigger()
}