package parameter

type ParameterIXI struct {
    intParameters     map[string]*IntParameter
    stringParameters  map[string]*StringParameter
    addIntHandlers    []IntParameterConsumer
    addStringHandlers []StringParameterConsumer
}

var globalInstance *ParameterIXI = nil

func IXI() *ParameterIXI {
    if globalInstance == nil {
        globalInstance = &ParameterIXI{
            make(map[string]*IntParameter),
            make(map[string]*StringParameter),
            make([]IntParameterConsumer, 0),
            make([]StringParameterConsumer, 0),
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

func (this *ParameterIXI) GetInts() map[string]*IntParameter {
    return this.intParameters
}

func (this *ParameterIXI) AddString(name string, defaultValue string, description string) *StringParameter {
    if this.intParameters[name] != nil {
        panic("duplicate parameter - \"" + name + "\" was defined already")
    }

    this.stringParameters[name] = newStringParameter(name, defaultValue, description)

    this.TriggerAddString(this.stringParameters[name])

    return this.stringParameters[name]
}

func (this *ParameterIXI) GetString(name string) *StringParameter {
    return this.stringParameters[name]
}

func (this *ParameterIXI) OnAddString(handler StringParameterConsumer) *ParameterIXI {
    this.addStringHandlers = append(this.addStringHandlers, handler)

    return this
}

func (this *ParameterIXI) TriggerAddString(param *StringParameter) {
    for _, handler := range this.addStringHandlers {
        handler(param)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////
