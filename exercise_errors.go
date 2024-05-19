package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
    z := x / 2
	iters := 0
    for ; math.Abs((z*z)-x) > 0.000000000001; {
        z -= (z*z - x) / (2*z)
		iters += 1
    }
	fmt.Printf("Number of iterations: %d\n", iters)
    return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}