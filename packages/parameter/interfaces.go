package parameter

type IntParameter interface {
    GetName() string
    GetValue() int
    GetValuePtr() *int
    GetDefaultValue() int
    GetDescription() string
}

type StringParameter interface {
    GetName() string
    GetValue() string
    GetValuePtr() *string
    GetDefaultValue() string
    GetDescription() string
}

type IntParameterConsumer func(param IntParameter)

type StringParameterConsumer func(param StringParameter)
