package config

type Config struct {
    sections map[string]*ConfigSection
}

func NewConfig() *Config {
    return &Config{
        sections: make(map[string]*ConfigSection),
    }
}

func (this *Config) AddSection(name string) *ConfigSection {
    section := this.GetSection(name)
    if section == nil {
        section = NewConfigSection(name)

        this.sections[name] = section
    }

    return section
}

func (this *Config) GetSection(name string) *ConfigSection {
    return this.sections[name]
}