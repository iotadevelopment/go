package tcp

import (
    "github.com/iotadevelopment/go/packages/network"
    "reflect"
)

//region errorConsumerEvent ////////////////////////////////////////////////////////////////////////////////////////////

func (this *errorConsumerEvent) Attach(callback ErrorConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *errorConsumerEvent) Detach(callback ErrorConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *errorConsumerEvent) Trigger(err error) {
    for _, callback := range this.callbacks {
        callback(err)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region peerConsumerEvent /////////////////////////////////////////////////////////////////////////////////////////////

func (this *peerConsumerEvent) Attach(callback PeerConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *peerConsumerEvent) Detach(callback PeerConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *peerConsumerEvent) Trigger(peer network.Peer) {
    for _, callback := range this.callbacks {
        callback(peer)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////