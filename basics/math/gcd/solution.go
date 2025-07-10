package main

import "fmt"

func minVal(condition bool, n1 int, n2 int) int {
	if condition {
		return n1
	} else {
		return n2
	}
}

func gcd(n1 int, n2 int) int {
	var min int = minVal(n1 < n2, n1, n2)
	var gcd int = 1
	for i := 2; i <= min; i++ {
		if n1%i == 0 && n2%i == 0 {
			gcd = i
		}
	}
	return gcd
}

func main() {
	var n1 int = 20
	var n2 int = 15
	var res int = gcd(n1, n2)
	fmt.Print(res)
}
