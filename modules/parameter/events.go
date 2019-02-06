package parameter

import "reflect"

//region intParameterEvent /////////////////////////////////////////////////////////////////////////////////////////////

func (this *intParameterEvent) Attach(callback IntParameterConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *intParameterEvent) Detach(callback IntParameterConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *intParameterEvent) Trigger(param *IntParameter) {
    for _, callback := range this.callbacks {
        callback(param)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region stringParameterEvent //////////////////////////////////////////////////////////////////////////////////////////

func (this *stringParameterEvent) Attach(callback StringParameterConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *stringParameterEvent) Detach(callback StringParameterConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *stringParameterEvent) Trigger(param *StringParameter) {
    for _, callback := range this.callbacks {
        callback(param)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////