package main

import (
    "core/crypto"
    "time"

    //"core/network/gossip"
    "core/ternary"
    "encoding/base64"
    "fmt"
    //"ixi"
    //ixigossip "ixi/gossip"
    //"iximodules/autopeering"
    //"iximodules/tcpserver"
    "os"
    //"strconv"
    //"time"
)

func main() {
    arguments := os.Args
    if len(arguments) == 1 {
        fmt.Println("Please provide a port number!")
        return
    }
    //port, _ := strconv.Atoi(arguments[1])

    fmt.Println("  _____ _____  _____   ___    ___")
    fmt.Println(" |_   _|  __ \\|_   _| |__ \\  / _ \\")
    fmt.Println("   | | | |__) | | |      ) || | | |")
    fmt.Println("   | | |  _  /  | |     / / | | | |")
    fmt.Println("  _| |_| | \\ \\ _| |_   / /_ | |_| |")
    fmt.Println(" |_____|_|  \\_\\_____| |____(_)___/")
    fmt.Println("")
    fmt.Println("Started successfully ...")
    fmt.Println("")

    // create IXI interface
    //IXI := ixi.NewIXI()

    sDec, _ := base64.StdEncoding.DecodeString("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")
    trits := ternary.BytesToTrits(sDec)[:8019]

    //sDec1, _ := base64.StdEncoding.DecodeString("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")
    //trits1 := ternary.BytesToTrits(sDec1)[:8019]

    go crypto.StartBackgroundWorker()

    collectedChannels := make([]chan ternary.Trinary, 5000)

    start1 := time.Now().UnixNano() / 1e6

    // issue hash requests
    for i := 0; i < 5000; i++ {
        resultChannel := crypto.FastHash(trits)

        collectedChannels[i] = resultChannel
    }

    // collect hash results
    for _, collectedChannel := range collectedChannels {
        <- collectedChannel
    }
    //fmt.Println(ternary.TritsToString(lastHash, 0, 243))
    now1 := time.Now().UnixNano() / 1e6
    fmt.Println(now1 - start1)

    //fmt.Println(ternary.TritsToString(hashedTrits, 0, 243))

    /*

    start1 := time.Now().UnixNano() / 1e6
    for i:= 0; i < 16; i++ {
        multiplexer := bcternary.NewBCTrinaryMultiplexer()
        multiplexer.Add(trits)
        multiplexer.Add(trits1)
        bcTrits, _ := multiplexer.Extract()

        multiCurl := crypto.NewBCCurl(243, 81)
        multiCurl.Absorb(bcTrits)
        hashedBcTrits := multiCurl.Squeeze(243)

        demux := bcternary.NewBCTrinaryDemultiplexer(hashedBcTrits)
        demux.Get(0)
        fmt.Println(ternary.TritsToString(demux.Get(1), 0, 243))

        //fmt.Println(ternary.TritsToString(demuxedTrits, 0, 243))
    }
    now1 := time.Now().UnixNano() / 1e6
    fmt.Println(now1 - start1)

    start2 := time.Now().UnixNano() / 1e6
    for i:= 0; i < 1000; i++ {
        crypto.RunHashCurl(trits)

        //fmt.Println(ternary.TritsToString(demuxedTrits, 0, 243))
    }
    now2 := time.Now().UnixNano() / 1e6
    fmt.Println(now2 - start2)



    // create a IXI module that dumps the tps count
/*
    cnt := 0
    start := time.Now().UnixNano() / 1e6
    IXI.Gossip.OnReceiveTransactionData(func(peer ixigossip.Peer, transactionData []byte) {
        cnt++
        now := time.Now().UnixNano() / 1e6

        if now - start > 1000 {
            fmt.Println(cnt, "TPS")

            start = now
            cnt = 0
        }
    })

    // load IXI modules
    tcpServer := tcpserver.NewTcpServer(IXI)
    peerManager := gossip.NewPeerManager(IXI)
    autoPeering := autopeering.NewAutoPeering(IXI, peerManager)

    // start the modules
    autoPeering.Start()
    tcpServer.Start(port)
*/
}
