package bcternary

import (
    "core/ternary"
    "errors"
    "strconv"
)

type BCTrinaryMultiplexer struct {
    trinaries []ternary.Trinary
}

func NewBCTrinaryMultiplexer() *BCTrinaryMultiplexer {
    this := &BCTrinaryMultiplexer{make([]ternary.Trinary, 0)}

    return this
}

func (this *BCTrinaryMultiplexer) Add(trinary ternary.Trinary) int {
    this.trinaries = append(this.trinaries, trinary)

    return len(this.trinaries) - 1
}

func (this *BCTrinaryMultiplexer) Get(index int) ternary.Trinary {
    return this.trinaries[index]
}

func (this *BCTrinaryMultiplexer) Extract() (*BCTrinary, error) {
    trinariesCount := len(this.trinaries)
    tritsCount := len(this.trinaries[0])

    result := &BCTrinary{
        Lo: make([]uint64, tritsCount),
        Hi: make([]uint64, tritsCount),
    }

    for i := 0; i < tritsCount; i++ {
        bcTrit := &BCTrit{0, 0}

        for j := 0; j < trinariesCount; j++ {
            switch this.trinaries[j][i] {
            case -1:
                bcTrit.Lo |= 1 << uint(j)

            case 1:
                bcTrit.Hi |= 1 << uint(j)

            case 0:
                bcTrit.Lo |= 1 << uint(j)
                bcTrit.Hi |= 1 << uint(j)

            default:
                return nil, errors.New("Invalid trit #" + strconv.Itoa(i) + " in trinary #" + strconv.Itoa(j))
            }
        }

        result.Lo[i] = bcTrit.Lo
        result.Hi[i] = bcTrit.Hi
    }

    return result, nil
}
