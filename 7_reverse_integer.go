package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var isNegative = false
	if x < 0 {
		isNegative = true
		x *= -1
	}
	var result = 0
	for tmp := x; tmp > 0; {
		result = result * 10 + tmp % 10
		tmp /= 10
	}
	if isNegative == true {
		result *= -1
	}
	if result> math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}

func main() {
	fmt.Print(reverse(12345))
}
