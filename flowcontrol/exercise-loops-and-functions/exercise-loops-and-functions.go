package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if z*z == x {
			fmt.Println("early break")
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
