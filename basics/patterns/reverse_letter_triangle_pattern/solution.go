package main

import "fmt"

func pattern(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Printf("%c ", 'A'+j)
		}
		fmt.Printf("\n")
	}
}
func main() {
	pattern(6)
}
