package main

import (
	"fmt"
)

// uint64 has a bigger positive range than int
func fib(n int) []uint64 {
	if n <= 0 {
		return []uint64{}
	}

	seq := make([]uint64, n+1)

	for i, _ := range seq {
		if i <= 1 {
			seq[i] = uint64(i)
		} else if i > 1 {
			seq[i] = seq[i-1] + seq[i-2]
		}
	}

	return seq
}

func main() {
	var input int

	fmt.Print("Enter a number: ")
	fmt.Scan(&input)

	fmt.Printf("Fibonnaci of %d is %v", input, fib(input))
}
