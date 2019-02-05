package config

import (
    "github.com/iotadevelopment/go/packages/ixi"
    "github.com/iotadevelopment/go/modules/parameter"
    "strings"
)

func configure() {
    parameter.Events.AddInt.Attach(func(param *parameter.IntParameter) {
        parsedAddress := strings.Split(param.Name, "/")
        if len(parsedAddress) != 2 {
            panic("invalid parameter name - expected format: \"<config_section>/<variable_name>\"")
        }

        // automatically create a config parameter
        section := AddSection(parsedAddress[0])
        section.AddIntValue(param.Value, parsedAddress[1], param.Description)
    })
}

func run() {}

var PLUGIN = ixi.NewPlugin(configure, run)
