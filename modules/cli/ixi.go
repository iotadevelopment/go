package cli

import (
    "flag"
    "fmt"
    "os"
    "path/filepath"
)

type CLIIXI struct {}

var globalInstance *CLIIXI = nil

func IXI() *CLIIXI {
    if globalInstance == nil {
        globalInstance = &CLIIXI{}
    }

    return globalInstance
}

func (this *CLIIXI) AddIntParameter(p *int, name string, usage string) {
    flag.IntVar(p, name, *p, usage)
}

func (this *CLIIXI) PrintUsage() {
    fmt.Fprintf(
        os.Stderr,
        "\n" +
        "GOIOTA 1.0\n\n" +
        "  A lightweight modular IOTA node.\n\n" +
        "Usage:\n\n" +
        "  %s [OPTIONS]\n\n" +
        "Options:\n\n", filepath.Base(os.Args[0]))

    flag.PrintDefaults()
}