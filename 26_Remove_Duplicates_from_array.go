package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	n := 0
	for i := 0; i < len(nums); i++ {
		if n == 0 || nums[i] != nums[n-1] {
			nums[n] = nums[i]
			n++
		}
	}
	return n
}

func main() {
	testArray1 := []int{1, 1, 2, 3, 4, 4, 6, 7}
	testArray2 := []int{1, 1, 2}
	testArray3 := []int{1, 1, 1}
	//fmt.Print(append(testArray1[:2], testArray1[3:]...))
	fmt.Println(removeDuplicates(testArray3))
	fmt.Println(testArray1)
	fmt.Println(removeDuplicates(testArray2))
	fmt.Println(testArray2)
}
