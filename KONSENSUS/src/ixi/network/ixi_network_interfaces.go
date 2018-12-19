package network

import "net"

type ErrorConsumer func(e error)

type SocketConsumer func(c net.Conn)

type SocketDataConsumer func(c net.Conn, data []byte)

type SocketErrorConsumer func(c net.Conn, err error)
