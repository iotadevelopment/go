package config

type ConfigIntValue struct {
    name string
    description string
    value *int
}

func NewConfigIntValue(p *int, name string, description string) *ConfigIntValue {
    return &ConfigIntValue{
        value: p,
        name: name,
        description: description,
    }
}

func (this *ConfigIntValue) GetValue() int {
    return *this.value
}

func (this *ConfigIntValue) SetValue(value int) {
    *this.value = value
}

func (this *ConfigIntValue) ValuePtr() *int {
    return this.value
}

func (this *ConfigIntValue) Description() string {
    return this.description
}