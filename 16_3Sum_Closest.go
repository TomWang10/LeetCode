package main

import (
	"fmt"
	"sort"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	min := math.MaxInt32
	closestValue := 0
	sort.Ints(nums)
	length := len(nums)
	for i :=0; i < length - 2; i++ {
		start := i + 1
		end := length - 1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			var tmpValue int
			// on right side
			if target > sum {
				tmpValue = target - sum
				if tmpValue < min {
					min = tmpValue
					closestValue = sum
					//for ;end - 1 > start && nums[end] == nums[end-1]; end--{}
					// process duplicates of sums[start]
					//for ;start + 1 < end && nums[start] == nums[start+1]; start++{}
				}
				start++
				// on left side
			} else if target < sum {
				tmpValue = sum - target
				if tmpValue < min {
					min = tmpValue
					closestValue = sum
					//for ;end - 1 > start && nums[end] == nums[end-1]; end--{}
					// process duplicates of sums[start]
					//for ;start + 1 < end && nums[start] == nums[start+1]; start++{}
				}
				end--
			} else {
				return target
			}
		}
		// process duplicates of sums[i]
		for ;i + 1 < length -2 && nums[i] == nums[i+1]; i++ {}
	}
	return closestValue
}

func main() {
	fmt.Println(threeSumClosest([]int{0,2,1,-3}, 1))
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{13,2,0,-14,-20,19,8,-5,-13,-3,20,15,20,5,13,14,-17,-7,12,-6,0,20,-19,-1,-15,-2,8,
	-2,-9,13,0,-3,-18,-9,-9,-19,17,-14,-19,-4,-16,2,0,9,5,-7,-4,20,18,9,0,12,-1,10,-17,-11,16,-13,-14,-3,0,2,-18,2,8,20,
	-15,3,-13,-12,-2,-19,11,11,-10,1,1,-10,-2,12,0,17,-19,-7,8,-19,-17,5,-5,-10,8,0,-12,4,19,2,0,12,14,-9,15,7,0,-16,-5,
	16,-12,0,2,-16,14,18,12,13,5,0,5,6}, -59))
}
