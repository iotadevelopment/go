package cli

import (
    "flag"
    "github.com/iotadevelopment/go/packages/ixi"
    "github.com/iotadevelopment/go/packages/parameter"
    "strings"
)

var MODULE = ixi.NewIXIModule().OnConfigure(func() {
    parameter.IXI().OnAddInt(func(param *parameter.IntParameter) {
        flagName := strings.Replace(strings.Replace(strings.ToLower(param.GetName()), "/", "-", 1), "_", "-", -1)

        IXI().AddIntParameter(param.GetValuePtr(), flagName, param.GetDescription())
    })

    flag.Usage = func() { IXI().PrintUsage() }
}).OnRun(func() {
    flag.Parse()
})