package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1) // from the hint on the tour site
	last := float64(.9)
	for math.Abs(z-last) > 0.000001 {
		last = z
		z = z - ((z*z - x) / (2 * z))
		fmt.Printf("last: %f z:%f diff:%f\n", last, z, (z - last))
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
