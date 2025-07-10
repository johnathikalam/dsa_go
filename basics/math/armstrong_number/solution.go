package main

import "fmt"

func isArmstrongNumber(n int) bool {
	var duplicate int = n
	var num int = 0
	for n > 0 {
		i := n % 10
		num += (i * i * i)
		n /= 10
	}
	return (duplicate == num)
}

func main() {
	var n int = 153
	var res bool = isArmstrongNumber(n)
	if res {
		fmt.Print("True")
	} else {
		fmt.Print("False")
	}
}
