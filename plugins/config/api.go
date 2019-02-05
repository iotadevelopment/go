package config

var config *Config = &Config{
    sections: make(map[string]*ConfigSection),
}

func AddSection(name string) *ConfigSection {
    return config.AddSection(name)
}

func GetSection(name string) *ConfigSection {
    return config.GetSection(name)
}