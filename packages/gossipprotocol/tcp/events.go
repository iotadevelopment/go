package tcp

import "reflect"

type protocolEvents struct {
    ReceivePortData               *dataEvent
    ReceiveTransactionData        *dataEvent
    ReceiveTransactionRequestData *dataEvent
    Error                         *errorEvent
}

//region dataEvent /////////////////////////////////////////////////////////////////////////////////////////////////////

type dataEvent struct {
    callbacks map[uintptr]DataConsumer
}

func (this *dataEvent) Attach(callback DataConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *dataEvent) Detach(callback DataConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *dataEvent) Trigger(data []byte) {
    for _, callback := range this.callbacks {
        callback(data)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region dataEvent /////////////////////////////////////////////////////////////////////////////////////////////////////

type errorEvent struct {
    callbacks map[uintptr]DataConsumer
}

func (this *errorEvent) Attach(callback DataConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *errorEvent) Detach(callback DataConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *errorEvent) Trigger(data []byte) {
    for _, callback := range this.callbacks {
        callback(data)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////
