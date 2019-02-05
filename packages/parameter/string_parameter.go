package parameter

type stringParameterImplementation struct {
    name         string
    value        *string
    defaultValue string
    description  string
}

func newStringParameter(name string, defaultValue string, description string) StringParameter {
    // create parameter
    parameter := &stringParameterImplementation{
        name:          name,
        defaultValue:  defaultValue,
        value:         &defaultValue,
        description:   description,
    }

    return parameter
}

func (this *stringParameterImplementation) GetName() string {
    return this.name
}

func (this *stringParameterImplementation) GetValue() string {
    return *this.value
}

func (this *stringParameterImplementation) GetValuePtr() *string {
    return this.value
}

func (this *stringParameterImplementation) GetDefaultValue() string {
    return this.defaultValue
}

func (this *stringParameterImplementation) GetDescription() string {
    return this.description
}

