package ixi

import (
    "bufio"
    "fmt"
    "os"
)

var loadedModules = make([]*IXIModule, 0)

func Load(modules ...*IXIModule) {
    loadedModules = append(loadedModules, modules...)

    for _, module := range modules {
        module.TriggerLoad()
    }
}

func Run() {
    for _, module := range loadedModules {
        module.TriggerConfigure()
    }

    for _, module := range loadedModules {
        module.TriggerRun()
    }

    fmt.Print("Press 'Enter' to exit...")
    bufio.NewReader(os.Stdin).ReadRune()
}
