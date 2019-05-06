// +build ignore

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

/**
 * Go's basic types are:
 *   - bool
 *   - string
 *   - int int8 int 16 int 32 int64
 *   - uint uint8 uint16 uint32 uint64 uintptr
 *   - byte // alias for uint8
 *   - rune // alias for int 32
 *          // represents a Unicode code point
 *   - float32 float64
 *   - complex64 complex128
}
 */
func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}