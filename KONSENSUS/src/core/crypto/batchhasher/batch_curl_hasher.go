package batchhasher

import (
    "core/crypto/bctcrypto"
    "core/crypto"
    "core/ternary/bcternary"
    "core/ternary"
    "fmt"
    "strconv"
    "time"
)

type CurlHashRequest struct {
    input ternary.Trinary
    output chan ternary.Trinary
}

type BatchCurlHasher struct {
    hashRequests chan CurlHashRequest
    hashLength int
    rounds int
}

func NewBatchCurlHasher(hashLength int, rounds int) *BatchCurlHasher {
    this := &BatchCurlHasher{
        hashLength: hashLength,
        rounds: rounds,
        hashRequests: make(chan CurlHashRequest),
    }

    go this.StartDispatcher()

    return this
}

func (this *BatchCurlHasher) StartDispatcher() {
    for {
        collectedHashRequests := make([]CurlHashRequest, 0)

        // wait for first request to start processing at all
        collectedHashRequests = append(collectedHashRequests, <- this.hashRequests)

        // collect additional requests that arrive within the timeout
    CollectAdditionalRequests:
        for {
            select {
            case hashRequest := <- this.hashRequests:
                collectedHashRequests = append(collectedHashRequests, hashRequest)

                if len(collectedHashRequests) == strconv.IntSize {
                    break CollectAdditionalRequests
                }
            case <- time.After(50 * time.Millisecond):
                break CollectAdditionalRequests
            }
        }

        go this.ProcessHashes(collectedHashRequests)
    }
}

func (this *BatchCurlHasher) ProcessHashes(collectedHashRequests []CurlHashRequest) {
    if len(collectedHashRequests) > 1 {
        // multiplex the requests
        multiplexer := bcternary.NewBCTrinaryMultiplexer()
        for _, hashRequest := range collectedHashRequests {
            multiplexer.Add(hashRequest.input)
        }
        bcTrinary, err := multiplexer.Extract()
        if err != nil {
            fmt.Println(err)
        }

        // calculate the hash
        bctCurl := bctcrypto.NewBCCurl(this.hashLength, this.rounds)
        bctCurl.Reset()
        bctCurl.Absorb(bcTrinary)

        // extract the results from the demultiplexer
        demux := bcternary.NewBCTrinaryDemultiplexer(bctCurl.Squeeze(243))
        for i, hashRequest := range collectedHashRequests {
            hashRequest.output <- demux.Get(i)
            close(hashRequest.output)
        }
    } else {
        var resp = make(ternary.Trinary, this.hashLength)

        curl := crypto.NewCurl(this.hashLength, this.rounds)
        curl.Absorb(collectedHashRequests[0].input, 0, len(collectedHashRequests[0].input))
        curl.Squeeze(resp, 0, this.hashLength)

        collectedHashRequests[0].output <- resp
        close(collectedHashRequests[0].output)
    }
}

func (this *BatchCurlHasher) Hash(trinary ternary.Trinary) chan ternary.Trinary {
    hashRequest := CurlHashRequest{
        input: trinary,
        output: make(chan ternary.Trinary, 1),
    }

    this.hashRequests <- hashRequest

    return hashRequest.output
}