package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev2, prev1 := 0, 1
	num := 0
	return func() int {
		num += 1
		if num == 1 {
			return 0
		}
		if num == 2 {
			return 1
		}
		next := prev1 + prev2
		prev2 = prev1
		prev1 = next
		return next
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
