package main

import "fmt"

func frequencies(arr []int, n int) {
	var hashMap = make(map[int]int)
	for i := 0; i < n; i++ {
		hashMap[arr[i]] += 1
	}

	var minIndex = 0
	var minCount = n
	var maxIndex = 0
	var maxCount = 0

	for key, value := range hashMap {
		if value > maxCount {
			maxIndex = key
			maxCount = value
		}
		if value < minCount {
			minIndex = key
			minCount = value
		}
	}
	fmt.Printf("%d %d\n", maxIndex, maxCount)
	fmt.Printf("%d %d\n", minIndex, minCount)
}

func main() {
	var arr = []int{10, 5, 10, 15, 10, 5}
	frequencies(arr, len(arr))
}
