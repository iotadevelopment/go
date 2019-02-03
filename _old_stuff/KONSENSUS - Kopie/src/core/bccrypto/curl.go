package bccrypto

import (
    "core/ternary"
)

var (
    curlHashers map[int]map[int]*Curl = make(map[int]map[int]*Curl)
)

type HashRequest struct {
    input ternary.Trinary
    output chan ternary.Trinary
}

type Curl struct {
    hashRequests chan HashRequest
}

func NewCurl() {

}