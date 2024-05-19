package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := x / 2
	iters := 0
    for ; math.Abs((z*z)-x) > 0.000000000001; {
        z -= (z*z - x) / (2*z)
		iters += 1
    }
	fmt.Printf("Number of iterations: %d\n", iters)
    return z
}

func main() {
    fmt.Println(Sqrt(2))
}