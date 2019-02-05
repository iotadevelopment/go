package config

func (this *ConfigSection) AddIntValue(p *int, name string, description string) *ConfigIntValue {
    this.intValues[name] = &ConfigIntValue{
        value: p,
        name: name,
        description: description,
    }

    return this.intValues[name]
}

func (this *ConfigSection) GetIntValue(name string) *ConfigIntValue {
    return this.intValues[name]
}