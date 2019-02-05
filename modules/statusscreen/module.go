package statusscreen

import (
    "fmt"
    "github.com/iotadevelopment/go/modules/gossip"
    "github.com/iotadevelopment/go/packages/ixi"
    "github.com/iotadevelopment/go/packages/transaction"
    "github.com/iotadevelopment/go/packages/network"
    "github.com/iotadevelopment/go/packages/terminal"
    "time"
)

var tpsReceived = 0
var tpsParsed = 0

func PrintBanner() {
    fmt.Println("")
    fmt.Println("   _____ _____  _____   ___    ___")
    fmt.Println("  |_   _|  __ \\|_   _| |__ \\  / _ \\")
    fmt.Println("    | | | |__) | | |      ) || | | |")
    fmt.Println("    | | |  _  /  | |     / / | | | |")
    fmt.Println("   _| |_| | \\ \\ _| |_   / /_ | |_| |")
    fmt.Println("  |_____|_|  \\_\\_____| |____(_)___/")
    fmt.Println("")
    fmt.Println("")
    fmt.Println("  Performance metrics")
    fmt.Println("  ===================")
    fmt.Println("")
}

func PrintTPSReceived() {
    fmt.Println("  Transactions Per Second (received): ", int(float64(tpsReceived) / REFRESH_INTERVAL.Seconds()))

    tpsReceived = 0
}

func PrintTPSParsed() {
    fmt.Println("  Transactions Per Second (parsed):   ", int(float64(tpsParsed) / REFRESH_INTERVAL.Seconds()))

    tpsParsed = 0
}

var MODULE = ixi.NewIXIModule().OnConfigure(func() {
    gossip.IXI().OnReceivePacketData(func(peer network.Peer, data []byte) {
        tpsReceived++
    }).OnReceiveTransaction(func(peer network.Peer, transaction transaction.Transaction) {
        tpsParsed++
    })
}).OnRun(func() {
    for {
        terminal.Clear()

        PrintBanner()
        PrintTPSReceived()
        PrintTPSParsed()

        time.Sleep(REFRESH_INTERVAL)
    }
})
