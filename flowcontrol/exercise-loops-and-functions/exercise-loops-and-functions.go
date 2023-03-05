package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		z -= (z*z - x) / (2 * z)
		if math.Abs(x-z*z) < 10e-14 {
			break
		}
	}
	return z
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println("i = ", i, "My Sqrt:", Sqrt(float64(i)), "math.Sqrt:", math.Sqrt(float64(i)))
	}
}
