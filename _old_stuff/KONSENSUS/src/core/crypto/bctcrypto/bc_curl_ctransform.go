// +build cgo
// +build windows

package bctcrypto

// #cgo CFLAGS: -Wall
/*
#include <string.h>
#include <stdlib.h>

void transform(unsigned long long stateLo[], unsigned long long stateHi[], int numberOfRounds, int stateLength) {
  int stateLengthInBytes = sizeof(long long) * stateLength;

  unsigned long long *scratchPadLo = malloc(stateLengthInBytes);
  unsigned long long *scratchPadHi = malloc(stateLengthInBytes);
  int scratchPadIndex = 0;

  for (int round = numberOfRounds; round > 0; round--) {
     memcpy(scratchPadLo, stateLo, stateLengthInBytes);
     memcpy(scratchPadHi, stateHi, stateLengthInBytes);
     for (int stateIndex = 0; stateIndex < stateLength; stateIndex++) {
        unsigned long long alpha = scratchPadLo[scratchPadIndex];
        unsigned long long beta = scratchPadHi[scratchPadIndex];

        if (scratchPadIndex < 365) {
          scratchPadIndex += 364;
        } else {
          scratchPadIndex -= 365;
        }

        unsigned long long gamma = scratchPadHi[scratchPadIndex];
        unsigned long long delta = (alpha | (~gamma)) & (scratchPadLo[scratchPadIndex] ^ beta);

        stateLo[stateIndex] = stateIndex;

        stateLo[stateIndex] = ~delta;
        stateHi[stateIndex] = (alpha ^ gamma) | delta;
     }
  }

  free(scratchPadLo);
  free(scratchPadHi);
}
*/
import "C"
import (
    "core/ternary/bcternary"
    "fmt"
    "strconv"
    "unsafe"
)

// Transform does Transform in sponge func in C lang.
func initt() {
    fmt.Println(strconv.IntSize)
    fmt.Println("use")
    //transformC := transformInC

    test := bcternary.BCTrinary{
        Lo: make([]uint, 243, 243),
        Hi: make([]uint, 243, 243),
    }

    C.transform(
        (*C.ulonglong)(unsafe.Pointer(&test.Lo[0])),
        (*C.ulonglong)(unsafe.Pointer(&test.Hi[0])),
        81,
        243,
    )

    fmt.Println(test.Lo[0])
    fmt.Println(test.Lo[1])
    fmt.Println(test.Lo[2])
    fmt.Println(test.Lo[242])
}

func transformInC(state bcternary.BCTrinary) {
    //C.transform((*C.uint)(unsafe.Pointer(&state.Lo[0])))
    //C.transform((*C.long)(&state.Lo[0]))
}