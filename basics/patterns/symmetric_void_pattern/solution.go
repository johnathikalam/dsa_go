package main

import "fmt"

func pattern(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Printf("*")
		}
		for j := 0; j < 2*i; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < n-i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("*")
		}
		for j := 0; j < 2*(n-i); j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func main() {
	pattern(6)
}
