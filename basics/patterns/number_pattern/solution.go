package main

import "fmt"

func pattern(n int) {
	for i := n; i >= 1; i-- {
		for j := n; j >= 1; j-- {
			if j <= i {
				fmt.Printf("%d ", i)
			} else {
				fmt.Printf("%d ", j)
			}
		}
		for j := 2; j <= n; j++ {
			if j <= i {
				fmt.Printf("%d ", i)
			} else {
				fmt.Printf("%d ", j)
			}
		}
		fmt.Print("\n")
	}
	for i := 1; i <= n; i++ {
		for j := n; j >= 1; j-- {
			if j <= i {
				fmt.Printf("%d ", i)
			} else {
				fmt.Printf("%d ", j)
			}
		}
		for j := 2; j <= n; j++ {
			if j <= i {
				fmt.Printf("%d ", i)
			} else {
				fmt.Printf("%d ", j)
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	pattern(6)
}
