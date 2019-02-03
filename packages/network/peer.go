package network

import (
    "bufio"
    "io"
    "net"
)

type peerImplementation struct {
    protocol            string
    conn                net.Conn
    receiveDataHandlers []DataConsumer
    disconnectHandlers  []Callback
    errorHandlers       []ErrorConsumer
}

func NewPeer(protocol string, conn net.Conn) Peer {
    this := &peerImplementation{
        protocol:            protocol,
        conn:                conn,
        receiveDataHandlers: make([]DataConsumer, 0),
        disconnectHandlers:  make([]Callback, 0),
        errorHandlers:       make([]ErrorConsumer, 0),
    }

    return this
}

func (this *peerImplementation) GetProtocol() string {
    return this.protocol
}

func (this *peerImplementation) GetConnection() net.Conn {
    return this.conn
}

func (this *peerImplementation) OnReceiveData(callback DataConsumer) Peer {
    this.receiveDataHandlers = append(this.receiveDataHandlers, callback)

    return this
}

func (this *peerImplementation) OnDisconnect(callback Callback) Peer {
    this.disconnectHandlers = append(this.disconnectHandlers, callback)

    return this
}

func (this *peerImplementation) OnError(callback ErrorConsumer) Peer {
    this.errorHandlers = append(this.errorHandlers, callback)

    return this
}

func (this *peerImplementation) TriggerReceiveData(data []byte) Peer {
    for _, receiveDataHandler := range this.receiveDataHandlers {
        receiveDataHandler(data)
    }

    return this
}

func (this *peerImplementation) TriggerDisconnect() Peer {
    for _, disconnectHandler := range this.disconnectHandlers {
        disconnectHandler()
    }

    return this
}

func (this *peerImplementation) TriggerError(err error) Peer {
    for _, errorHandler := range this.errorHandlers {
        errorHandler(err)
    }

    return this
}

func (this *peerImplementation) HandleConnection() {
    defer this.conn.Close()
    defer this.TriggerDisconnect()

    receiveBuffer := make([]byte, READ_BUFFER_SIZE)
    for {
        byteCount, err := bufio.NewReader(this.conn).Read(receiveBuffer)
        if err != nil {
            if err != io.EOF {
                this.TriggerError(err)
            }

            return
        }

        receivedData := make([]byte, byteCount)
        copy(receivedData, receiveBuffer)

        this.TriggerReceiveData(receivedData)
    }
}