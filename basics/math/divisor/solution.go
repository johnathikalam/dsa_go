package main

import (
	"fmt"
	"math"
	"sort"
)

func divisor(n int) []int {
	var divisors []int
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}
	return divisors
}

func main() {
	var n int = 36
	divisors := divisor(n)
	sort.Ints(divisors)
	fmt.Print(divisors)
}
