package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %s", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	z := x / 4

	// In use to debug number of steps
	//for i := 0; x - z * z < 1 / 10; i++ {
	for x-z*z < 1/10 {

		// In use to debug number of steps
		// print("step ", i, "\n")
		z -= (z*z - x) / (2 * z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(15.0))
	fmt.Println(Sqrt(-15.0))

}
