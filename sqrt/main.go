package main

import (
	"fmt"
	"math"
)

// Sqrt Calculate square root using Newton's method
func Sqrt(x float64) float64 {
	z := float64(1)
	var zFormer = 0.0
	for zFormer == 0 || math.Abs(z-zFormer) > 0.001 {
		if z != 1 {
			zFormer = z
		}
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(81))
}
