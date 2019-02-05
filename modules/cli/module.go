package cli

import (
    "flag"
    "github.com/iotadevelopment/go/packages/ixi"
    "github.com/iotadevelopment/go/packages/parameter"
    "strings"
)

func onAddInt(param *parameter.IntParameter) {
    flagName := strings.Replace(strings.Replace(strings.ToLower(param.GetName()), "/", "-", 1), "_", "-", -1)

    IXI().AddIntParameter(param.GetValuePtr(), flagName, param.GetDescription())
}

var MODULE = ixi.NewIXIModule().OnConfigure(func() {
    for _, param := range parameter.IXI().GetInts() {
        onAddInt(param)
    }

    parameter.IXI().OnAddInt(onAddInt)

    flag.Usage = func() { IXI().PrintUsage() }
}).OnRun(func() {
    flag.Parse()
})