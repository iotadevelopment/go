package cli

import (
    "flag"
    "github.com/iotadevelopment/go/modules/parameter"
    "github.com/iotadevelopment/go/packages/ixi"
    "strings"
)

func onAddIntParameter(param *parameter.IntParameter) {
    flagName := strings.Replace(strings.Replace(strings.ToLower(param.Name), "/", "-", 1), "_", "-", -1)

    addIntParameter(param.Value, flagName, param.Description)
}

func onAddStringParameter(param *parameter.StringParameter) {
    flagName := strings.Replace(strings.Replace(strings.ToLower(param.Name), "/", "-", 1), "_", "-", -1)

    addStringParameter(param.Value, flagName, param.Description)
}

func configure() {
    for _, param := range parameter.GetInts() {
        onAddIntParameter(param)
    }

    for _, param := range parameter.GetStrings() {
        onAddStringParameter(param)
    }

    parameter.Events.AddInt.Attach(onAddIntParameter)
    parameter.Events.AddString.Attach(onAddStringParameter)

    flag.Usage = printUsage
}

var PLUGIN = ixi.NewPlugin(configure, flag.Parse)