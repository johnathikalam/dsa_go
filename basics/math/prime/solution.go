package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int = 36
	res := isPrime(n)
	if res {
		fmt.Print("True")
	} else {
		fmt.Print("False")
	}
}
