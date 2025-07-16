package main

import "fmt"

func bubble_sort(arr []int, n int) {
	var flag bool = true
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j+1] < arr[j] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
				flag = false
			}
		}
		if flag {
			break
		}
	}
}

func main() {
	var arr = []int{3, 6, 4, 1, 7, 2, 9, 8}
	n := len(arr)
	bubble_sort(arr, n)
	fmt.Print(arr)
}
