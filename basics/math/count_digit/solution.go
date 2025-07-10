package main

import (
	"fmt"
	"math"
)

func countDigitWhile(n int) int {
	var cnt int = 0
	for n > 0 {
		cnt += 1
		n /= 10
	}
	return cnt
}

func countDigitLog(n int) int {
	return int(math.Log10(float64(n)) + 1)
}

func main() {
	var n int = 12345
	res := countDigitLog(n)
	fmt.Print(res)
}
