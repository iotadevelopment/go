package ixi

import "reflect"

type pluginEvents struct {
    Configure *event
    Run       *event
}

type event struct {
    callbacks map[uintptr]Callback
}

func (this *event) Attach(callback Callback) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *event) Detach(callback Callback) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *event) Trigger() {
    for _, callback := range this.callbacks {
        callback()
    }
}
