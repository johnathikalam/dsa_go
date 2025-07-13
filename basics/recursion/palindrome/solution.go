package main

import "fmt"

func isPalindrome(arr []int, start int, end int) bool {
	if start >= end {
		return true
	}
	if arr[start] != arr[end] {
		return false
	}
	return isPalindrome(arr, start+1, end-1)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	if isPalindrome(arr, 0, len(arr)-1) {
		fmt.Print("True")
	} else {
		fmt.Print("False")
	}
}
