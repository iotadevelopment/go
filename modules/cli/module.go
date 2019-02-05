package cli

import (
    "flag"
    "github.com/iotadevelopment/go/packages/ixi"
    "github.com/iotadevelopment/go/packages/parameter"
    "strings"
)

func onAddInt(param parameter.IntParameter) {
    flagName := strings.Replace(strings.Replace(strings.ToLower(param.GetName()), "/", "-", 1), "_", "-", -1)

    IXI().AddIntParameter(param.GetValuePtr(), flagName, param.GetDescription())
}

var PLUGIN = ixi.NewPlugin(func() {
    for _, param := range parameter.GetInts() {
        onAddInt(param)
    }

    parameter.Events.AddInt.Attach(onAddInt)

    flag.Usage = func() { IXI().PrintUsage() }
}, func() {
    flag.Parse()
})