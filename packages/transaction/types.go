package transaction

import "github.com/iotadevelopment/go/packages/ternary"

type Transaction struct {
    SignatureMessageFragment      ternary.Trits
    Address                       ternary.Trits
    Value                         int64
    ObsoleteTag                   ternary.Trits
    Timestamp                     uint64
    CurrentIndex                  uint64
    LatestIndex                   uint64
    BundleHash                    ternary.Trits
    TrunkTransactionHash          ternary.Trits
    BranchTransactionHash         ternary.Trits
    Tag                           ternary.Trits
    AttachmentTimestamp           uint64
    AttachmentTimestampLowerBound uint64
    AttachmentTimestampUpperBound uint64
    Nonce                         ternary.Trits

    Hash                          ternary.Trits
    WeightMagnitude               int
    Bytes                         []byte
    Trits                         ternary.Trits
}
