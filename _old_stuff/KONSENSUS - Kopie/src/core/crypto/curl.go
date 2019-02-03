package crypto

import (
    "core/ternary"
    "fmt"
    "math"
)

const (
    NUMBER_OF_ROUNDSP27 = 27
    NUMBER_OF_ROUNDSP81 = 81
    HASH_LENGTH         = 243
    STATE_LENGTH        = 3 * HASH_LENGTH
)

var (
    TRUTH_TABLE = ternary.Trinary{1, 0, -1, 2, 1, -1, 0, 2, -1, 1, 0}
)


type Hash interface {
    Initialize()
    InitializeCurl(trits *[]int8, length int, rounds int)
    Reset()
    Absorb(trits *[]int8, offset int, length int)
    Squeeze(resp []int8, offset int, length int) []int
}


type Curl struct {
    Hash
    state  ternary.Trinary
    rounds int
}

func (curl *Curl) Initialize() {
    curl.InitializeCurl(nil, 0, curl.rounds)
}

func (curl *Curl) InitializeCurl(trinary ternary.Trinary, length int, rounds int) {
    curl.rounds = rounds
    if trinary != nil {
        curl.state = trinary
    } else {
        curl.state = make(ternary.Trinary, STATE_LENGTH)
    }
}

func (curl *Curl) Reset() {
    curl.InitializeCurl(nil, 0, curl.rounds)
}

func (curl *Curl) Absorb(trinary ternary.Trinary, offset int, length int) {
    for {
        limit := int(math.Min(HASH_LENGTH, float64(length)))
        copy(curl.state, trinary[offset:offset+limit])
        curl.Transform()
        offset += HASH_LENGTH
        length -= HASH_LENGTH
        if length <= 0 {
            break
        }

        if false {
            fmt.Println(trinary)
            fmt.Println(limit)
        }
    }
}

func (curl *Curl) Squeeze(resp ternary.Trinary, offset int, length int) ternary.Trinary {
    for {
        limit := int(math.Min(HASH_LENGTH, float64(length)))
        copy(resp[offset:offset+limit], curl.state)
        curl.Transform()
        offset += HASH_LENGTH
        length -= HASH_LENGTH
        if length <= 0 {
            break
        }
    }
    return resp
}

func (curl *Curl) Transform() {
    var index = 0
    for round := 0; round < curl.rounds; round++ {
        stateCopy := make(ternary.Trinary, STATE_LENGTH)
        copy(stateCopy, curl.state)
        for i := 0; i < STATE_LENGTH; i++ {
            incr := 364
            if index >= 365 {
                incr = -365
            }
            index2 := index + incr
            curl.state[i] = TRUTH_TABLE[stateCopy[index]+(stateCopy[index2]<<2)+5]
            index = index2
        }
    }
}

func RunHashCurl(trits ternary.Trinary) ternary.Trinary {
    var curl = new(Curl)
    var resp = make(ternary.Trinary, HASH_LENGTH)

    curl.InitializeCurl(nil, 0, NUMBER_OF_ROUNDSP81)
    curl.Absorb(trits, 0, len(trits))
    curl.Squeeze(resp, 0, HASH_LENGTH)

    return resp
}
