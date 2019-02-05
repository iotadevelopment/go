package parameter

var (
    intParameters     = make(map[string]IntParameter)
    stringParameters  = make(map[string]StringParameter)
)

//region IXI METHODS ///////////////////////////////////////////////////////////////////////////////////////////////////

var Events = struct {
    AddInt    *intParameterEventManager
    AddString *stringParameterEventManager
}{
    AddInt:    &intParameterEventManager{make(map[uintptr]IntParameterConsumer)},
    AddString: &stringParameterEventManager{make(map[uintptr]StringParameterConsumer)},
}

func AddInt(name string, defaultValue int, description string) IntParameter {
    if intParameters[name] != nil {
        panic("duplicate parameter - \"" + name + "\" was defined already")
    }

    newParameter := newIntParameter(name, defaultValue, description)
    intParameters[name] = newParameter

    Events.AddInt.Trigger(newParameter)

    return newParameter
}

func GetInt(name string) IntParameter {
    return intParameters[name]
}

func GetInts() map[string]IntParameter {
    return intParameters
}

func AddString(name string, defaultValue string, description string) StringParameter {
    if intParameters[name] != nil {
        panic("duplicate parameter - \"" + name + "\" was defined already")
    }

    newParameter := newStringParameter(name, defaultValue, description)
    stringParameters[name] = newParameter

    Events.AddString.Trigger(newParameter)

    return stringParameters[name]
}

func GetString(name string) StringParameter {
    return stringParameters[name]
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////
