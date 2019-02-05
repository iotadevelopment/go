package config

type Config struct {
    sections map[string]*ConfigSection
}

type ConfigSection struct {
    name      string
    intValues map[string]*ConfigIntValue
}

type ConfigIntValue struct {
    name string
    description string
    value *int
}
