package config

type ConfigSection struct {
    name      string
    intValues map[string]*ConfigIntValue
}

func NewConfigSection(name string) *ConfigSection {
    return &ConfigSection{name, make(map[string]*ConfigIntValue)}
}

func (this *ConfigSection) AddIntValue(p *int, name string, description string) *ConfigIntValue {
    this.intValues[name] = NewConfigIntValue(p, name, description)

    return this.intValues[name]
}

func (this *ConfigSection) GetIntValue(name string) *ConfigIntValue {
    return this.intValues[name]
}