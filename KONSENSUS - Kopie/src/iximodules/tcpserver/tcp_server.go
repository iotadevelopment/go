package tcpserver

import (
    "bufio"
    "fmt"
    "io"
    "ixi"
    "net"
    "strconv"
)

const (
    READ_BUFFER_SIZE = 81920
)

type TcpServer struct {
    ixi *ixi.IXI
}

func NewTcpServer(ixi *ixi.IXI) *TcpServer {
    return &TcpServer{ixi}
}

func (this *TcpServer) Start(port int) {
    l, err := net.Listen("tcp4", "127.0.0.1:"+strconv.Itoa(port))
    if err != nil {
        this.ixi.Network.TriggerError(err)

        return
    }
    defer l.Close()

    for {
        c, err := l.Accept()
        if err != nil {
            fmt.Println(err)

            return
        }

        go this.handleConnection(c)
    }
}

func (this *TcpServer) handleConnection(c net.Conn) {
    defer c.Close()
    defer this.ixi.Network.TriggerClientDisconnect(c)

    this.ixi.Network.TriggerClientConnect(c)

    receiveBuffer := make([]byte, READ_BUFFER_SIZE)
    for {
        byteCount, err := bufio.NewReader(c).Read(receiveBuffer)
        if err != nil {
            if err != io.EOF {
                this.ixi.Network.TriggerClientError(c, err)
            }

            return
        }

        receivedData := make([]byte, byteCount)
        copy(receivedData, receiveBuffer)

        err = this.ixi.Network.TriggerClientReceiveData(c, receivedData)
        if err != nil {
            this.ixi.Network.TriggerClientError(c, err)

            return
        }
    }
}
