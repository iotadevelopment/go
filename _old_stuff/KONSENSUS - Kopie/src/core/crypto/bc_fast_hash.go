package crypto

import (
    "core/bcternary"
    "core/ternary"
    "fmt"
    "time"
)

type HashRequest struct {
    input ternary.Trinary
    output chan ternary.Trinary
}

var (
    hashRequests = make(chan HashRequest)
)

func StartBackgroundWorker() {
    bcCurl := NewBCCurl(243, 81)

    for {
        collectedHashRequests := make([]HashRequest, 0)

        // wait for first request to start processing at all
        collectedHashRequests = append(collectedHashRequests, <- hashRequests)

        // collect additional requests that arrive within the timeout
        CollectAdditionalRequests:
        for {
            select {
            case hashRequest := <- hashRequests:
                collectedHashRequests = append(collectedHashRequests, hashRequest)

                if len(collectedHashRequests) == 64 {
                    break CollectAdditionalRequests
                }
            case <- time.After(50 * time.Millisecond):
                break CollectAdditionalRequests
            }
        }

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
            bcCurl.Reset()
            bcCurl.Absorb(bcTrinary)

            // extract the results from the demultiplexer
            demux := bcternary.NewBCTrinaryDemultiplexer(bcCurl.Squeeze(243))
            for i, hashRequest := range collectedHashRequests {
                hashRequest.output <- demux.Get(i)
                close(hashRequest.output)
            }
        } else {

        }
    }
}

func FastHash(trinary ternary.Trinary) chan ternary.Trinary {
    hashRequest := HashRequest{
        input: trinary,
        output: make(chan ternary.Trinary, 1),
    }

    hashRequests <- hashRequest

    return hashRequest.output
}
