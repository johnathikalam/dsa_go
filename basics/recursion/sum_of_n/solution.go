package main

import "fmt"

func sumOfN(n int, sum int) {
	if n == 0 {
		fmt.Println(sum)
		return
	}
	sumOfN(n-1, sum+n)
}

func main() {
	var n int = 5
	sumOfN(n, 0)
}
