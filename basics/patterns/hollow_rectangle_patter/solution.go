package main

import "fmt"

func pattern(n int) {
	for i := 0; i < n; i++ {
		if i == 0 || i == n-1 {
			for j := 0; j < n; j++ {
				fmt.Print("*")
			}
			fmt.Print("\n")
		} else {
			fmt.Print("*")
			for j := 1; j < n-1; j++ {
				fmt.Print(" ")
			}
			fmt.Print("*\n")
		}
	}
}

func main() {
	pattern(6)
}
