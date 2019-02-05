package parameter

type intParameterImplementation struct {
    name         string
    value        *int
    defaultValue int
    description  string
}

func newIntParameter(name string, defaultValue int, description string) IntParameter {
    parameter := &intParameterImplementation{
        name:          name,
        defaultValue:  defaultValue,
        value:         &defaultValue,
        description:   description,
    }

    return parameter
}

func (this *intParameterImplementation) GetName() string {
    return this.name
}

func (this *intParameterImplementation) GetValue() int {
    return *this.value
}

func (this *intParameterImplementation) GetValuePtr() *int {
    return this.value
}

func (this *intParameterImplementation) GetDefaultValue() int {
    return this.defaultValue
}

func (this *intParameterImplementation) GetDescription() string {
    return this.description
}
