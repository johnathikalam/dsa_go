package main

import "fmt"

func fibonacci(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int = 10
	fmt.Print(fibonacci(n))
}
