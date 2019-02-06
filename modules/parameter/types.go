package parameter

type moduleEvents struct {
    AddInt    *intParameterEvent
    AddString *stringParameterEvent
}

type intParameterEvent struct {
    callbacks map[uintptr]IntParameterConsumer
}

type stringParameterEvent struct {
    callbacks map[uintptr]StringParameterConsumer
}

type IntParameter struct {
    Name         string
    Value        *int
    DefaultValue int
    Description  string
}

type StringParameter struct {
    Name         string
    Value        *string
    DefaultValue string
    Description  string
}

type IntParameterConsumer = func(param *IntParameter)

type StringParameterConsumer = func(param *StringParameter)
