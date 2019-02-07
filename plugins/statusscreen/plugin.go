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
}

func PrintPerformanceMetrics() {
    fmt.Println("")
    fmt.Println("  Performance metrics")
    fmt.Println("  ===================")
    fmt.Println("")
    fmt.Println("  Transactions Per Second (received): ", int(float64(tpsReceived) / REFRESH_INTERVAL.Seconds()))
    fmt.Println("  Transactions Per Second (parsed):   ", int(float64(tpsParsed) / REFRESH_INTERVAL.Seconds()))

    tpsReceived = 0
    tpsParsed = 0
}

func PrintNeighbours() {
    neighborManager := gossip.GetNeighborManager()

    fmt.Println("")
    fmt.Println("  Gossip Neighbors")
    fmt.Println("  ================")
    fmt.Println("")

    neighbors := make([]string, 0)
    for _, neighbor := range neighborManager.GetStaticNeighbors() {
        neighbors = append(neighbors, neighbor.GetAddress())
    }

    for _, neighbor := range neighborManager.GetDynamicNeighbors() {
        line := neighbor.GetAddress() + " (dynamic)"

        if neighbor.UpstreamConnected() && !neighbor.DownstreamConnected() {
            line += " [warning: only upstream connected]"
        }

        if neighbor.DownstreamConnected() && !neighbor.UpstreamConnected() {
            line += " [warning: only downstream connected]"
        }

        neighbors = append(neighbors, line)
    }

    if (len(neighbors) >= 1) {
        for _, line := range neighbors {
            fmt.Println("  " + line)
        }
    } else {
        fmt.Println("  <None>")
    }
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
        PrintPerformanceMetrics()
        PrintNeighbours()

        time.Sleep(REFRESH_INTERVAL)
    }
})
