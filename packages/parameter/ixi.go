package parameter

type ParameterIXI struct {
    intParameters  map[string]*IntParameter
    addIntHandlers []IntParameterConsumer
}

var globalInstance *ParameterIXI = nil

func IXI() *ParameterIXI {
    if globalInstance == nil {
        globalInstance = &ParameterIXI{
            make(map[string]*IntParameter),
            make([]IntParameterConsumer, 0),
        }
    }

    return globalInstance
}

//region IXI METHODS ///////////////////////////////////////////////////////////////////////////////////////////////////

func (this *ParameterIXI) AddInt(name string, defaultValue int, description string) *IntParameter {
    if this.intParameters[name] != nil {
        panic("duplicate parameter - \"" + name + "\" was defined already")
    }

    this.intParameters[name] = newIntParameter(name, defaultValue, description)

    this.TriggerAddInt(this.intParameters[name])

    return this.intParameters[name]
}

func (this *ParameterIXI) GetInt(name string) *IntParameter {
    return this.intParameters[name]
}

func (this *ParameterIXI) OnAddInt(handler IntParameterConsumer) *ParameterIXI {
    this.addIntHandlers = append(this.addIntHandlers, handler)

    return this
}

func (this *ParameterIXI) TriggerAddInt(param *IntParameter) {
    for _, handler := range this.addIntHandlers {
        handler(param)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////
