package main

import "fmt"

func isPalindrome(n int) bool {
	var rev int = 0
	var duplicate int = n
	for n > 0 {
		rev = (rev * 10) + (n % 10)
		n /= 10
	}
	return (duplicate == rev)
}
func main() {
	var n int = 123454321
	var res bool = isPalindrome(n)
	if res {
		fmt.Print("Palindrome Number")
	} else {
		fmt.Print("Not Palindrome")
	}
}
