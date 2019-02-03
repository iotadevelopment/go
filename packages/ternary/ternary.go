package ternary

// a Trit can have the values 0, 1 and -1
type Trit int8

// a Trinary consists out of many Trits
type Trits []Trit

func (this Trits) TrailingZeroes() int {
    zeros := 0
    index := len(this) - 1
    for this[index] == 0 {
        zeros++

        index--
    }

    return zeros
}

func (this Trits) ToInt64() int64 {
    var val int64
    for i := len(this) - 1; i >= 0; i-- {
        val = val * 3 + int64(this[i])
    }

    return val
}

func (this Trits) ToUint64() uint64 {
    var val uint64
    for i := len(this) - 1; i >= 0; i-- {
        val = val * 3 + uint64(this[i])
    }

    return val
}

func (this Trits) ToString() string {
    return TritsToString(this, 0, len(this))
}