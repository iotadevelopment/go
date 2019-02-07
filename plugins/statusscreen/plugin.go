package statusscreen

import (
    "fmt"
    "github.com/iotadevelopment/go/packages/ixi"
    "github.com/iotadevelopment/go/packages/terminal"
    "github.com/iotadevelopment/go/packages/transaction"
    "github.com/iotadevelopment/go/plugins/gossip"
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

var PLUGIN = ixi.NewPlugin(func() {
    neighborManager := gossip.GetNeighborManager()

    neighborManager.Events.AddNeighbor.Attach(func(neighbor *gossip.Neighbor) {
        neighbor.Events.ReceiveTransactionData.Attach(func(data []byte) {
            tpsReceived++
        })

        neighbor.Events.ReceiveTransaction.Attach(func(transaction *transaction.Transaction) {
            tpsParsed++
        })
    })
}, func() {
    for {
        terminal.Clear()

        PrintBanner()
        PrintTPSReceived()
        PrintTPSParsed()

        time.Sleep(REFRESH_INTERVAL)
    }
})
