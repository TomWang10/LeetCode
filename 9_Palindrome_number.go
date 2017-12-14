package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	length := 1
	for x /length >= 10 {
		length *= 10
	}

	for x > 0 {
		if x /length != x % 10 {
			return false
		}
		x = (x % length) / 10
		length /= 100
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(123321))
	fmt.Println(isPalindrome(-2147483648))
	fmt.Println(isPalindrome(-2147447412))
}
