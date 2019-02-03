package config

type ConfigIXI struct {
    config *Config
}

var globalInstance *ConfigIXI = nil

func IXI() *ConfigIXI {
    if globalInstance == nil {
        globalInstance = &ConfigIXI{
            config: NewConfig(),
        }
    }

    return globalInstance
}

func (this *ConfigIXI) AddSection(name string) *ConfigSection {
    return this.config.AddSection(name)
}

func (this *ConfigIXI) GetSection(name string) *ConfigSection {
    return this.config.GetSection(name)
}