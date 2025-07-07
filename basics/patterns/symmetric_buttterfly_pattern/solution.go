package main

import "fmt"

func pattern(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		for j := 0; j < 2*(n-i-1); j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print("*")
		}
		for j := 0; j < 2*(i+1); j++ {
			fmt.Print(" ")
		}
		for j := 0; j < n-i-1; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}

func main() {
	pattern(6)
}
