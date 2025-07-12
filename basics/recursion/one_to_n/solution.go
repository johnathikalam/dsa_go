package main

import "fmt"

func oneToN(n int) {
	if n <= 0 {
		return
	}
	oneToN(n - 1)
	fmt.Printf("%d ", n)
}

func nToOne(n int) {
	if n <= 0 {
		return
	}
	fmt.Printf("%d ", n)
	nToOne(n - 1)
}

func main() {
	n := 5
	oneToN(n)
	fmt.Println()
	nToOne(n)
}
