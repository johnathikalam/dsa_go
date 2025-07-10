package main

import "fmt"

func reverse(n int) int {
	var rev int = 0
	for n > 0 {
		rev = (rev * 10) + (n % 10)
		n /= 10
	}
	return rev
}

func main() {
	var n int = 12345
	var res int = reverse(n)
	fmt.Print(res)
}
