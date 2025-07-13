package main

import "fmt"

func reverse(arr []int, start int, end int) {
	if start > end {
		return
	}
	temp := arr[start]
	arr[start] = arr[end]
	arr[end] = temp
	reverse(arr, start+1, end-1)
}

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(arr, 0, len(arr)-1)
	fmt.Print(arr)
}
