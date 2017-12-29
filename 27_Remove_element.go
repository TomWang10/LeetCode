package main

import (
	"fmt"
	"sort"
)

func removeElement(nums []int, val int) int {
	sort.Ints(nums)
	left := 0
	right := len(nums)-1
	oldLength := len(nums)
	for left <= right {
		mid := nums[(left + right)/2]
		if val == mid {
			// find start position
			start := (left + right)/2
			for ; start - 1 >= 0 && nums[start - 1] == val; start--{}
			// find end position
			end := (left + right)/2
			for ; end + 1 < len(nums) && nums[end + 1] == val; end++ {}
			nums = append(nums[:start], nums[end+1:]...)
			return oldLength - (end - start + 1)
		} else if val > mid {
			left = (left + right)/2 + 1
		} else {
			right = (left + right)/2 - 1
		}
	}
	return len(nums)
}

func main() {
	fmt.Println(removeElement([]int{3,2,2,3}, 3))
	fmt.Println(removeElement([]int{3,2,1,4,4,5,2,3}, 1))
	fmt.Println(removeElement([]int{3,2,1,1,4,5,4,5,2,3}, 1))
	fmt.Println(removeElement([]int{}, 0))
	fmt.Println(removeElement([]int{2}, 3))
	fmt.Println(removeElement([]int{1}, 1))
	fmt.Println(removeElement([]int{4,5}, 4))
}
