package recording

import (
    "github.com/iotadevelopment/go/plugins/gossip"
    "os"
)

var receivedData = make(chan []byte)

func collectData(data []byte) {
    receivedData <- data
}

func writeData() {
    if f, err := os.OpenFile("tangle.rec", os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666); err != nil {
        panic(err)
    } else {
        defer f.Close()

        for msg := range receivedData {
            f.Write(msg)
        }
    }
}

func setupNodeEventHandlers(neighbor *gossip.Neighbor) {
    neighbor.Events.ReceiveData.Attach(collectData)
}

func tearNodeDownEventHandlers(neighbor *gossip.Neighbor) {
    neighbor.Events.ReceiveData.Detach(collectData)
}
