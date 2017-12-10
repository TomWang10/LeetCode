package main

import (
	"fmt"
	"bytes"
)

func convert(s string, numRows int) string {
	var resultArray = make([][]byte, numRows, numRows)
	var n, i = len(s), 0
	for i < n {
		for j := 0; j < numRows && i < n; j++{
			resultArray[j] = append(resultArray[j], s[i])
			i++
		}
		for j := numRows - 2; j >=1 && i < n; j-- {
			resultArray[j] = append(resultArray[j], s[i])
			i++
		}
	}
	//var result string
	var result bytes.Buffer
	for _, v := range resultArray {
		result.Write(v)
	}
	return result.String()
}

func main() {
	fmt.Print(convert("PAYPALISHIRING", 3))
}