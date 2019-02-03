package parameter

type IntParameter struct {
    name         string
    value        *int
    defaultValue int
    description  string
}

func newIntParameter(name string, defaultValue int, description string) *IntParameter {
    // create parameter
    parameter := &IntParameter{
        name:          name,
        defaultValue:  defaultValue,
        value:         &defaultValue,
        description:   description,
    }

    return parameter
}

func (this *IntParameter) GetName() string {
    return this.name
}

func (this *IntParameter) GetValue() int {
    return *this.value
}

func (this *IntParameter) GetValuePtr() *int {
    return this.value
}

func (this *IntParameter) GetDefaultValue() int {
    return this.defaultValue
}

func (this *IntParameter) GetDescription() string {
    return this.description
}
