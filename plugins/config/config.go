package config

func (this *Config) AddSection(name string) *ConfigSection {
    section := this.GetSection(name)
    if section == nil {
        section = &ConfigSection{
            name,
            make(map[string]*ConfigIntValue),
        }

        this.sections[name] = section
    }

    return section
}

func (this *Config) GetSection(name string) *ConfigSection {
    return this.sections[name]
}