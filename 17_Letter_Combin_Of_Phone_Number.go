package main

import (
	"fmt"
)

var charTable = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func Search(digits string, index int, subString string, result *[]string) {

	if index == len(digits) {
		*result = append(*result, subString)
		return
	}
	singleDigit := digits[index]
	letter := charTable[singleDigit - '0']
	for _, v := range letter {
		Search(digits, index+1, subString + string(v), result)
	}
}

func letterCombinations(digits string) []string {
	var result []string
	if len(digits) == 0 {
		return result
	}

	Search(digits, 0, "", &result)
	return result
}

func main() {
	//fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("3"))
}
