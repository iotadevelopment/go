package transaction

import "github.com/iotadevelopment/go/packages/ternary"

type Transaction interface {
    GetHash() ternary.Trits
    GetBytes() []byte

    SetBytes(bytes []byte)
}
