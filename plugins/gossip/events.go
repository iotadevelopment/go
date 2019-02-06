package gossip

import (
    "github.com/iotadevelopment/go/packages/network"
    "github.com/iotadevelopment/go/packages/transaction"
    "reflect"
)

type gossipEvents struct {
    Connect            *peerEvent
    Error              *errorEvent
    ReceiveData        *peerDataEvent
    Disconnect         *peerEvent
    PeerError          *peerErrorEvent
    ReceivePacketData  *peerDataEvent
    ReceiveTransaction *peerTransactionEvent
}

//region peerEvent /////////////////////////////////////////////////////////////////////////////////////////////////////

type peerEvent struct {
    callbacks map[uintptr]PeerConsumer
}

func (this *peerEvent) Attach(callback PeerConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *peerEvent) Detach(callback PeerConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *peerEvent) Trigger(peer network.Peer) {
    for _, callback := range this.callbacks {
        callback(peer)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region errorEvent ////////////////////////////////////////////////////////////////////////////////////////////////////

type errorEvent struct {
    callbacks map[uintptr]ErrorConsumer
}

func (this *errorEvent) Attach(callback ErrorConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *errorEvent) Detach(callback ErrorConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *errorEvent) Trigger(err error) {
    for _, callback := range this.callbacks {
        callback(err)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region peerDataEvent /////////////////////////////////////////////////////////////////////////////////////////////////

type peerDataEvent struct {
    callbacks map[uintptr]PeerDataConsumer
}

func (this *peerDataEvent) Attach(callback PeerDataConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *peerDataEvent) Detach(callback PeerDataConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *peerDataEvent) Trigger(peer network.Peer, data []byte) {
    for _, callback := range this.callbacks {
        callback(peer, data)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region peerErrorEvent ////////////////////////////////////////////////////////////////////////////////////////////////

type peerErrorEvent struct {
    callbacks map[uintptr]PeerErrorConsumer
}

func (this *peerErrorEvent) Attach(callback PeerErrorConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *peerErrorEvent) Detach(callback PeerErrorConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *peerErrorEvent) Trigger(peer network.Peer, err error) {
    for _, callback := range this.callbacks {
        callback(peer, err)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region peerTransactionEvent //////////////////////////////////////////////////////////////////////////////////////////

type peerTransactionEvent struct {
    callbacks map[uintptr]PeerTransactionConsumer
}

func (this *peerTransactionEvent) Attach(callback PeerTransactionConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *peerTransactionEvent) Detach(callback PeerTransactionConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *peerTransactionEvent) Trigger(peer network.Peer, transaction *transaction.Transaction) {
    for _, callback := range this.callbacks {
        callback(peer, transaction)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////
