package model

import (
    "core/crypto"
    "fmt"
    "core/ternary"
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

func NewTransactionFromBytes(bytes *[]byte) {
    trits := ternary.BytesToTrits(bytes)[:8019]
    hashTrits := crypto.RunHashCurl(&trits)

    if false {
        fmt.Println(trits)
        fmt.Println(hashTrits)
    }

    //
    //fmt.Println(ternary.TritsToString(&hashTrits, 0, len(hashTrits)))
}