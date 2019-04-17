package tcpprotocol

import "reflect"

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
