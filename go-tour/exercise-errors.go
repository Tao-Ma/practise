package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

// Take care of recursively call Error() string.
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	var z float64 = 1
	delta := 0.00001
	times := 0
	if x < 0.0 {
		return 0, ErrNegativeSqrt(x)
	}

	for math.Abs(z-math.Sqrt(x)) > delta {
		if times++; times > 100 {
			break
		}
		z = z - (z*z-x)/(2*z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
