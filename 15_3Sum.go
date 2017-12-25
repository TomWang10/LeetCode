package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	numsLength := len(nums)
	for i := 0; i < numsLength; i++ {
		target := -nums[i]
		left := i+1
		right := numsLength-1
		for left < right {
			sum := nums[right] + nums[left]
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				//found
				result = append(result, []int{nums[i], nums[right], nums[left]})
				// process duplicate 3
				for tmp := nums[right]; tmp == nums[right] && right > left; right-- {}
				// process duplicate 2
				for tmp := nums[left]; tmp == nums[left] && right > left; left++ {}
			}
		}
		for ;i + 1 < numsLength-1 && nums[i + 1] == nums[i]; i++ {}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
