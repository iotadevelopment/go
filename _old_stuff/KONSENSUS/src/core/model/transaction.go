package model

import (
    "fmt"
    "github.com/iotadevelopment/go/packages/curl"
    "github.com/iotadevelopment/go/packages/ternary"
)

type Transaction struct {
    Hash                     []byte
    TrunkTransaction         []byte
    BranchTransaction        []byte
    Bundle                   []byte
    Address                  []byte
    Value                    int64
    Timestamp                int
    TXTimestamp              int
    AttachmentTimestamp      int
    CurrentIndex             int
    SignatureMessageFragment []int
}

func NewTransaction() *Transaction {
    this := &Transaction{}

    return this
}

func NewTransactionFromBytes(bytes []byte) {
    trits := ternary.BytesToTrits(bytes)[:8019]
    hashTrits := <- curl.CURLP81.Hash(trits)

    if false {
        fmt.Println(trits)
        fmt.Println(hashTrits)
    }

    //
    //fmt.Println(ternary.TritsToString(&hashTrits, 0, len(hashTrits)))
}