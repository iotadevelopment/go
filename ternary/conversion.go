package ternary

func BytesToTrits(bytes []byte) Trinary {
	size := len(bytes)
	trits := make([]Trit, size*NUMBER_OF_TRITS_IN_A_BYTE)

	for i := 0; i < size; i++ {
		v := int(bytes[i])
		if int8(bytes[i]) < 0 {
			v -= 13
		}

		for j := 0; j < NUMBER_OF_TRITS_IN_A_BYTE; j++ {
			trits[i*NUMBER_OF_TRITS_IN_A_BYTE+j] = BYTES_TO_TRITS[v*NUMBER_OF_TRITS_IN_A_BYTE+j]
		}
	}

	return trits
}
