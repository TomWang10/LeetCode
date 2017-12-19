package main

import (
	"fmt"
)

func maxArea(height []int) int {
	maxContainer := 0
	arrLen := len(height)
	for k, v := range height {
		for i := k + 1; i < arrLen; i++ {
			vertical := 0
			if v > height[i] {
				vertical = height[i]
			} else {
				vertical = v
			}
			tmpArea := vertical * (i - k)
			if maxContainer < tmpArea {
				maxContainer = tmpArea
			}
		}

	}
	return maxContainer
}

func main() {
	heightArray := []int{3, 1, 6, 4}
	fmt.Println(maxArea(heightArray))
}
