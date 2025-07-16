package main

import "fmt"

func selection_sort(arr []int, n int) {
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		temp := arr[min]
		arr[min] = arr[i]
		arr[i] = temp
	}
}

func main() {
	var arr = []int{3, 6, 4, 1, 7, 2, 9, 8}
	n := len(arr)
	selection_sort(arr, n)
	fmt.Print(arr)
}
