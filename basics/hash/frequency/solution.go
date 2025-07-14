package main

import "fmt"

func frequency(arr []int, n int) {
	var hashMap = make(map[int]int)
	for i := 0; i < n; i++ {
		hashMap[arr[i]] += 1
	}
	for key, value := range hashMap {
		fmt.Printf("%d  %d\n", key, value)
	}

}

func main() {
	var arr = []int{10, 5, 10, 15, 10, 5}
	frequency(arr, len(arr))
}
