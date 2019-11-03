package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x/4
	for i := 0; x - z * z < 1 / 10; i++ {
		print("step ", i, "\n")
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(123847238946))
}
