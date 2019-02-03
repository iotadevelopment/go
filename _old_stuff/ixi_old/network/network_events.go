package network

import "net"

func (this *IXINetwork) OnError(eventHandler ErrorConsumer) {
    this.errorHandlers = append(this.errorHandlers, eventHandler)
}

func (this *IXINetwork) OnClientConnect(eventHandler SocketConsumer) {
    this.clientConnectHandlers = append(this.clientConnectHandlers, eventHandler)
}

func (this *IXINetwork) OnClientDisconnect(eventHandler SocketConsumer) {
    this.clientDisconnectHandlers = append(this.clientDisconnectHandlers, eventHandler)
}

func (this *IXINetwork) OnClientReceiveData(eventHandler SocketDataConsumer) {
    this.clientReceiveDataHandlers = append(this.clientReceiveDataHandlers, eventHandler)
}

func (this *IXINetwork) OnClientError(eventHandler SocketErrorConsumer) {
    this.clientErrorHandlers = append(this.clientErrorHandlers, eventHandler)
}

func (this *IXINetwork) TriggerError(err error) {
    for _, handler := range this.errorHandlers {
        handler(err)
    }
}

func (this *IXINetwork) TriggerClientConnect(c net.Conn) {
    for _, onConnectHandler := range this.clientConnectHandlers {
        onConnectHandler(c)
    }
}

func (this *IXINetwork) TriggerClientDisconnect(c net.Conn) {
    for _, handler := range this.clientDisconnectHandlers {
        handler(c)
    }
}

func (this *IXINetwork) TriggerClientReceiveData(c net.Conn, data []byte) error {
    for _, handler := range this.clientReceiveDataHandlers {
        handler(c, data)
    }

    return nil
}

func (this *IXINetwork) TriggerClientError(c net.Conn, err error) {
    for _, handler := range this.clientErrorHandlers {
        handler(c, err)
    }
}