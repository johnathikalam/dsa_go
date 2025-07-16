package main

import "fmt"

func insertion_sort(arr []int, n int) {
	for i := 0; i < n; i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			temp := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = temp
			j--
		}
	}
}

func main() {
	var arr = []int{3, 6, 4, 1, 7, 2, 9, 8}
	n := len(arr)
	insertion_sort(arr, n)
	fmt.Print(arr)
}
