package main

import "fmt"

func pattern(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%c ", 'A'+n-i+j-1)
		}
		fmt.Printf("\n")
	}
}

func main() {
	pattern(6)
}
