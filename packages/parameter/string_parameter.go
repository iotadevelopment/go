package parameter

type StringParameter struct {
    name         string
    value        *string
    defaultValue string
    description  string
}

func newStringParameter(name string, defaultValue string, description string) *StringParameter {
    // create parameter
    parameter := &StringParameter{
        name:          name,
        defaultValue:  defaultValue,
        value:         &defaultValue,
        description:   description,
    }

    return parameter
}

func (this *StringParameter) GetName() string {
    return this.name
}

func (this *StringParameter) GetValue() string {
    return *this.value
}

func (this *StringParameter) GetValuePtr() *string {
    return this.value
}

func (this *StringParameter) GetDefaultValue() string {
    return this.defaultValue
}

func (this *StringParameter) GetDescription() string {
    return this.description
}

