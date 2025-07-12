package main

import "fmt"

func factorial(n int, fact int) {
	if n == 1 {
		fmt.Println(fact)
		return
	}
	factorial(n-1, fact*n)
}

func main() {
	var n int = 3
	factorial(n, 1)
}
