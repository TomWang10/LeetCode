package main

import (
	"fmt"
	"strings"
	"math"
)

func myAtoi(str string) int {
	if str == "" {
		return 0
	}
	str = strings.Trim(str, " ")
	var negativeFlag int32 = 1
	var result int32 = 0
	isOverflow := false
	for k, v := range str {
		if v == '-' && k == 0 {
			negativeFlag = -1
		} else if v == '+' && k == 0 {
			negativeFlag = 1
		} else {
			if v > 57 || v < 48 {
				break
			}
			oldResult := result
			result = result * 10
			if result / 10 != oldResult {
				isOverflow = true
				break
			}
			result += int32(v) - 48
			//result = result * 10 + (int32(v) - 48)
			if result < 0 {
				isOverflow = true
				break
			}
		}
	}
	if isOverflow {
		if negativeFlag == 1 {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}
	}
	result *= negativeFlag
	if result > math.MaxInt32 {
		result = math.MaxInt32
	} else if result < math.MinInt32 {
		result = math.MinInt32
	}
	return int(result)
}

func main() {
	fmt.Println(myAtoi("              010"))
	fmt.Println(myAtoi("+-2"))
	fmt.Println(myAtoi(""))
	fmt.Println(myAtoi("1"))
	fmt.Println(myAtoi("-12341"))
	fmt.Println(myAtoi("+12341"))
	fmt.Println(myAtoi("  -0012a42"))
	fmt.Println(myAtoi("2147483648"))
	fmt.Println(myAtoi("9223372036854775809"))
	fmt.Println(myAtoi("-9223372036854775809"))
	fmt.Println(myAtoi("    10522545459"))
}
