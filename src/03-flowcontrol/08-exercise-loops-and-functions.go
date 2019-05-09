// +build ignore

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x/2

	for i := 0; i < 10; i++ {
		z -= (z * z - x) / (2.0 * z)
	}

	return z
}

func main() {
	for i := 0.0; i < 10; i++ {
		expected := math.Sqrt(i)
		actual := Sqrt(i)
		fmt.Println(expected, actual, expected - actual)
	}
}