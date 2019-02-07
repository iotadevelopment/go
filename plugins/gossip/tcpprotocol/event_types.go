package tcpprotocol

import "reflect"

//region intEvent //////////////////////////////////////////////////////////////////////////////////////////////////////

type intEvent struct {
    callbacks map[uintptr]IntConsumer
}

func (this *intEvent) Attach(callback IntConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *intEvent) Detach(callback IntConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *intEvent) Trigger(number int) {
    for _, callback := range this.callbacks {
        callback(number)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

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
