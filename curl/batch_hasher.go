package curl

import (
    "fmt"
    "github.com/iotadevelopment/go/ternary"
    "strconv"
    "time"
)

type HashRequest struct {
    input ternary.Trinary
    output chan ternary.Trinary
}

type BatchHasher struct {
    hashRequests chan HashRequest
    hashLength int
    rounds int
}

func NewBatchHasher(hashLength int, rounds int) *BatchHasher {
    this := &BatchHasher{
        hashLength: hashLength,
        rounds: rounds,
        hashRequests: make(chan HashRequest),
    }

    go this.StartDispatcher()

    return this
}

func (this *BatchHasher) StartDispatcher() {
    for {
        collectedHashRequests := make([]HashRequest, 0)

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

func (this *BatchHasher) ProcessHashes(collectedHashRequests []HashRequest) {
    if len(collectedHashRequests) > 1 {
        // multiplex the requests
        multiplexer := ternary.NewBCTernaryMultiplexer()
        for _, hashRequest := range collectedHashRequests {
            multiplexer.Add(hashRequest.input)
        }
        bcTrinary, err := multiplexer.Extract()
        if err != nil {
            fmt.Println(err)
        }

        // calculate the hash
        bctCurl := NewBCTCurl(this.hashLength, this.rounds)
        bctCurl.Reset()
        bctCurl.Absorb(bcTrinary)

        // extract the results from the demultiplexer
        demux := ternary.NewBCTernaryDemultiplexer(bctCurl.Squeeze(243))
        for i, hashRequest := range collectedHashRequests {
            hashRequest.output <- demux.Get(i)
            close(hashRequest.output)
        }
    } else {
        var resp = make(ternary.Trinary, this.hashLength)

        curl := NewCurl(this.hashLength, this.rounds)
        curl.Absorb(collectedHashRequests[0].input, 0, len(collectedHashRequests[0].input))
        curl.Squeeze(resp, 0, this.hashLength)

        collectedHashRequests[0].output <- resp
        close(collectedHashRequests[0].output)
    }
}

func (this *BatchHasher) Hash(trinary ternary.Trinary) chan ternary.Trinary {
    hashRequest := HashRequest{
        input: trinary,
        output: make(chan ternary.Trinary, 1),
    }

    this.hashRequests <- hashRequest

    return hashRequest.output
}