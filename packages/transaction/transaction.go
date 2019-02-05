package transaction

import (
    "github.com/iotadevelopment/go/packages/curl"
    "github.com/iotadevelopment/go/packages/ternary"
)

type transactionImplementation struct {
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

func FromTrits(trits ternary.Trits, optionalHash ...ternary.Trits) Transaction {
    hash := <- curl.CURLP81.Hash(trits)

    transaction := &transactionImplementation{
        SignatureMessageFragment:      trits[SIGNATURE_MESSAGE_FRAGMENT_OFFSET:SIGNATURE_MESSAGE_FRAGMENT_END],
        Address:                       trits[ADDRESS_OFFSET:ADDRESS_END],
        Value:                         trits[VALUE_OFFSET:VALUE_END].ToInt64(),
        ObsoleteTag:                   trits[OBSOLETE_TAG_OFFSET:OBSOLETE_TAG_END],
        Timestamp:                     trits[TIMESTAMP_OFFSET:TIMESTAMP_END].ToUint64(),
        CurrentIndex:                  trits[CURRENT_INDEX_OFFSET:CURRENT_INDEX_END].ToUint64(),
        LatestIndex:                   trits[LATEST_INDEX_OFFSET:LATEST_INDEX_END].ToUint64(),
        BundleHash:                    trits[BUNDLE_HASH_OFFSET:BUNDLE_HASH_END],
        TrunkTransactionHash:          trits[TRUNK_TRANSACTION_HASH_OFFSET:TRUNK_TRANSACTION_HASH_END],
        BranchTransactionHash:         trits[BRANCH_TRANSACTION_HASH_OFFSET:BRANCH_TRANSACTION_HASH_END],
        Tag:                           trits[TAG_OFFSET:TAG_END],
        AttachmentTimestamp:           trits[ATTACHMENT_TIMESTAMP_OFFSET:ATTACHMENT_TIMESTAMP_END].ToUint64(),
        AttachmentTimestampLowerBound: trits[ATTACHMENT_TIMESTAMP_LOWER_BOUND_OFFSET:ATTACHMENT_TIMESTAMP_LOWER_BOUND_END].ToUint64(),
        AttachmentTimestampUpperBound: trits[ATTACHMENT_TIMESTAMP_UPPER_BOUND_OFFSET:ATTACHMENT_TIMESTAMP_UPPER_BOUND_END].ToUint64(),
        Nonce:                         trits[NONCE_OFFSET:NONCE_END],

        Hash:                          hash,
        WeightMagnitude:               hash.TrailingZeroes(),
        Trits:                         trits,
    }

    return transaction
}

func FromBytes(bytes []byte) Transaction {
    transaction := FromTrits(ternary.BytesToTrits(bytes)[:TRANSACTION_SIZE])
    transaction.SetBytes(bytes)

    return transaction
}

func (this *transactionImplementation) GetHash() ternary.Trits {
    return this.Hash
}

func (this *transactionImplementation) SetBytes(bytes []byte) {
    this.Bytes = bytes
}

func (this *transactionImplementation) GetBytes() []byte {
    return this.Bytes
}