package main

import (
	"fmt"
	"math"
)

var gCoinArray = []int{1, 3, 5}

func GetMinCoinCount(c int) int {
	if 0 == c {
		return 0
	}
	minCount := math.MaxInt32
	for _, v := range gCoinArray {
		if v <= c {
			tmpCount := GetMinCoinCount(c - v) + 1
			if tmpCount < minCount {
				minCount = tmpCount
			}
		}
	}
	return minCount
}

func main() {
	fmt.Println(GetMinCoinCount(12))
}
