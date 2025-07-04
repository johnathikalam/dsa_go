package main

import "fmt"

func pattern(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d", j+1)
		}
		for j := 0; j < 2*(n-i)-2; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j <= i; j++ {

			fmt.Printf("%d", i-j+1)
		}
		fmt.Printf("\n")
	}
}

func main() {
	pattern(6)
}
