package model

import (
    "github.com/iotadevelopment/go/packages/curl"
    "github.com/iotadevelopment/go/packages/ternary"
)

const (
    // sizes of the transaction fields
    SIGNATURE_MESSAGE_FRAGMENT_SIZE         = 6561
    ADDRESS_SIZE                            = 243
    VALUE_SIZE                              = 81
    OBSOLETE_TAG_SIZE                       = 81
    TIMESTAMP_SIZE                          = 27
    CURRENT_INDEX_SIZE                      = 27
    LATEST_INDEX_SIZE                       = 27
    BUNDLE_HASH_SIZE                        = 243
    TRUNK_TRANSACTION_HASH_SIZE             = 243
    BRANCH_TRANSACTION_HASH_SIZE            = 243
    TAG_SIZE                                = 81
    ATTACHMENT_TIMESTAMP_SIZE               = 27
    ATTACHMENT_TIMESTAMP_LOWER_BOUND_SIZE   = 27
    ATTACHMENT_TIMESTAMP_UPPER_BOUND_SIZE   = 27
    NONCE_SIZE                              = 81

    // offsets of the transaction fields
    SIGNATURE_MESSAGE_FRAGMENT_OFFSET       = 0
    ADDRESS_OFFSET                          = SIGNATURE_MESSAGE_FRAGMENT_END
    VALUE_OFFSET                            = ADDRESS_END
    OBSOLETE_TAG_OFFSET                     = VALUE_END
    TIMESTAMP_OFFSET                        = OBSOLETE_TAG_END
    CURRENT_INDEX_OFFSET                    = TIMESTAMP_END
    LATEST_INDEX_OFFSET                     = CURRENT_INDEX_END
    BUNDLE_HASH_OFFSET                      = LATEST_INDEX_END
    TRUNK_TRANSACTION_HASH_OFFSET           = BUNDLE_HASH_END
    BRANCH_TRANSACTION_HASH_OFFSET          = TRUNK_TRANSACTION_HASH_END
    TAG_OFFSET                              = BRANCH_TRANSACTION_HASH_END
    ATTACHMENT_TIMESTAMP_OFFSET             = TAG_END
    ATTACHMENT_TIMESTAMP_LOWER_BOUND_OFFSET = ATTACHMENT_TIMESTAMP_END
    ATTACHMENT_TIMESTAMP_UPPER_BOUND_OFFSET = ATTACHMENT_TIMESTAMP_LOWER_BOUND_END
    NONCE_OFFSET                            = ATTACHMENT_TIMESTAMP_UPPER_BOUND_END

    // ends of the transaction fields
    SIGNATURE_MESSAGE_FRAGMENT_END          = SIGNATURE_MESSAGE_FRAGMENT_OFFSET + SIGNATURE_MESSAGE_FRAGMENT_SIZE
    ADDRESS_END                             = ADDRESS_OFFSET + ADDRESS_SIZE
    VALUE_END                               = VALUE_OFFSET + VALUE_SIZE
    OBSOLETE_TAG_END                        = OBSOLETE_TAG_OFFSET + OBSOLETE_TAG_SIZE
    TIMESTAMP_END                           = TIMESTAMP_OFFSET + TIMESTAMP_SIZE
    CURRENT_INDEX_END                       = CURRENT_INDEX_OFFSET + CURRENT_INDEX_SIZE
    LATEST_INDEX_END                        = LATEST_INDEX_OFFSET + LATEST_INDEX_SIZE
    BUNDLE_HASH_END                         = BUNDLE_HASH_OFFSET + BUNDLE_HASH_SIZE
    TRUNK_TRANSACTION_HASH_END              = TRUNK_TRANSACTION_HASH_OFFSET + TRUNK_TRANSACTION_HASH_SIZE
    BRANCH_TRANSACTION_HASH_END             = BRANCH_TRANSACTION_HASH_OFFSET + BRANCH_TRANSACTION_HASH_SIZE
    TAG_END                                 = TAG_OFFSET + TAG_SIZE
    ATTACHMENT_TIMESTAMP_END                = ATTACHMENT_TIMESTAMP_OFFSET + ATTACHMENT_TIMESTAMP_SIZE
    ATTACHMENT_TIMESTAMP_LOWER_BOUND_END    = ATTACHMENT_TIMESTAMP_LOWER_BOUND_OFFSET + ATTACHMENT_TIMESTAMP_LOWER_BOUND_SIZE
    ATTACHMENT_TIMESTAMP_UPPER_BOUND_END    = ATTACHMENT_TIMESTAMP_UPPER_BOUND_OFFSET + ATTACHMENT_TIMESTAMP_UPPER_BOUND_SIZE
    NONCE_END                               = NONCE_OFFSET + NONCE_SIZE

    // the full size of a transaction
    TRANSACTION_SIZE                        = NONCE_END
)

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
}

func NewTransactionFromTrits(trits ternary.Trits, optionalHash ...ternary.Trits) *Transaction {
    hash := <- curl.CURLP81.Hash(trits)

    transaction := &Transaction{
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
    }

    return transaction
}

func NewTransactionFromBytes(bytes []byte) *Transaction {
    return NewTransactionFromTrits(ternary.BytesToTrits(bytes)[:TRANSACTION_SIZE])

    //fmt.Println(ternary.TritsToString(hashTrits, 0, len(hashTrits)))
}
