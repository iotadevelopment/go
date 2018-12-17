package ternary

import "bytes"

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

func TritsToString(trits Trinary, offset int, size int) string {
	var buffer bytes.Buffer
	for i := 0; i < (size + NUMBER_OF_TRITS_IN_A_TRYTE - 1) / NUMBER_OF_TRITS_IN_A_TRYTE; i++ {
		j := int(trits[offset + i * NUMBER_OF_TRITS_IN_A_TRYTE]) + int(trits[offset + i * NUMBER_OF_TRITS_IN_A_TRYTE + 1]) * NUMBER_OF_TRITS_IN_A_TRYTE + int(trits[offset + i * NUMBER_OF_TRITS_IN_A_TRYTE + 2]) * NUMBER_OF_TRITS_IN_A_TRYTE * NUMBER_OF_TRITS_IN_A_TRYTE;
		if j < 0 {
			j += len(TRYTE_ALPHABET)
		}
		buffer.WriteString(TRYTE_ALPHABET[j]);
	}

	return buffer.String()
}