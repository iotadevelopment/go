package parameter

import "reflect"

//region intParameterEventManager //////////////////////////////////////////////////////////////////////////////////////

type intParameterEventManager struct {
    callbacks map[uintptr]IntParameterConsumer
}

func (this *intParameterEventManager) Attach(callback IntParameterConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *intParameterEventManager) Detach(callback IntParameterConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *intParameterEventManager) Trigger(param IntParameter) {
    for _, callback := range this.callbacks {
        callback(param)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region stringParameterEventManager ///////////////////////////////////////////////////////////////////////////////////

type stringParameterEventManager struct {
    callbacks map[uintptr]StringParameterConsumer
}

func (this *stringParameterEventManager) Attach(callback StringParameterConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *stringParameterEventManager) Detach(callback StringParameterConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *stringParameterEventManager) Trigger(param StringParameter) {
    for _, callback := range this.callbacks {
        callback(param)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////