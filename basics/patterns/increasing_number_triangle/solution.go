package main

import "fmt"

func pattern(n int) {
	cnt := 1
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", cnt)
			cnt += 1
		}
		fmt.Print("\n")
	}

}

func main() {
	pattern(6)
}
