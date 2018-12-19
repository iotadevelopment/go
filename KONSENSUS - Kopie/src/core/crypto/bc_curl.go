package crypto

import (
    "core/bcternary"
    "core/ternary"
)

const (
    HIGH_LONG_BITS = 0xFFFFFFFFFFFFFFFF
)

type BCCurl struct {
    hashLength int
    numberOfRounds int
    stateLength int
    state bcternary.BCTrinary
    cTransform func()
}

func NewBCCurl(hashLength int, numberOfRounds int) *BCCurl {
    this := &BCCurl{
        hashLength: hashLength,
        numberOfRounds: numberOfRounds,
        stateLength: ternary.NUMBER_OF_TRITS_IN_A_TRYTE * hashLength,
        state: bcternary.BCTrinary{
            Lo: make([]uint64, ternary.NUMBER_OF_TRITS_IN_A_TRYTE * hashLength),
            Hi: make([]uint64, ternary.NUMBER_OF_TRITS_IN_A_TRYTE * hashLength),
        },
        cTransform: nil,
    }

    this.Reset()

    return this
}

func (this *BCCurl) Reset() {
    for i:= 0; i < this.stateLength; i++ {
        this.state[i] = bcternary.BCTrit{HIGH_LONG_BITS, HIGH_LONG_BITS}
    }
}

func (this *BCCurl) Transform() {
    scratchPad := make(bcternary.BCTrinary, this.stateLength)
    scratchPadIndex := 0

    for round := this.numberOfRounds; round > 0; round-- {
        copy(scratchPad, this.state)
        for stateIndex := 0; stateIndex < this.stateLength; stateIndex++ {
            alpha := scratchPad[scratchPadIndex].Lo
            beta := scratchPad[scratchPadIndex].Hi

            if scratchPadIndex < 365 {
                scratchPadIndex += 364
            } else {
                scratchPadIndex -= 365
            }

            gamma := scratchPad[scratchPadIndex].Hi
            delta := (alpha | (^gamma)) & (scratchPad[scratchPadIndex].Lo ^ beta)

            this.state[stateIndex] = bcternary.BCTrit{^delta, (alpha ^ gamma) | delta}
        }
    }
}

func (this *BCCurl) Absorb(bcTrits bcternary.BCTrinary) {
    length := len(bcTrits)
    offset := 0

    for {
        var lengthToCopy int
        if length < this.hashLength {
            lengthToCopy = length
        } else {
            lengthToCopy = this.hashLength
        }

        copy(this.state[0:lengthToCopy], bcTrits[offset:offset + lengthToCopy])
        this.Transform()

        offset += lengthToCopy
        length -= lengthToCopy

        if length <= 0 {
            break
        }
    }
}

func (this *BCCurl) Squeeze(tritCount int) bcternary.BCTrinary {
    result := make(bcternary.BCTrinary, tritCount)
    hashCount := tritCount / this.hashLength

    for i := 0; i < hashCount; i++ {
        copy(result[i * this.hashLength : (i + 1) * this.hashLength], this.state[0 : this.hashLength])
        this.Transform()
    }

    last := tritCount - hashCount * this.hashLength
    copy(result[tritCount - last : ], this.state[0 : last])
    if tritCount % this.hashLength != 0 {
        this.Transform()
    }

    return result
}