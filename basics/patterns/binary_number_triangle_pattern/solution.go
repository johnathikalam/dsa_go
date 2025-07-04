package main

import "fmt"

func pattern(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if (i+j)%2 == 0 {
				fmt.Printf("1")
			} else {
				fmt.Printf("0")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	pattern(6)
}
