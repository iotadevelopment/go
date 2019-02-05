package config

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