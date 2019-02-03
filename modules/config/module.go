package config

import (
    "github.com/iotadevelopment/go/packages/ixi"
    "github.com/iotadevelopment/go/packages/parameter"
    "strings"
)

var MODULE = ixi.NewIXIModule().OnConfigure(func() {
    parameter.IXI().OnAddInt(func(param *parameter.IntParameter) {
        parsedAddress := strings.Split(param.GetName(), "/")
        if len(parsedAddress) != 2 {
            panic("invalid parameter name - expected format: \"<config_section>/<variable_name>\"")
        }

        // automatically create a config parameter
        section := IXI().AddSection(parsedAddress[0])
        section.AddIntValue(param.GetValuePtr(), parsedAddress[1], param.GetDescription())
    })
})
