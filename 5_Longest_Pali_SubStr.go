package main

import (
	"fmt"
	"bytes"
)

func longestPalindrome(s string) string {
	var strBuffer bytes.Buffer
	strBuffer.WriteRune('#')
	for _, v := range s {
		strBuffer.WriteRune(v)
		strBuffer.WriteRune('#')
	}
	maxRight, maxPos := 0, 0
	radiusTable := make([]int, len(strBuffer.String()))
	strBufferLen := len(strBuffer.String())
	tmpStr := strBuffer.String()
	for i := 1; i < strBufferLen; i++ {
		if i < maxRight {
			radiusTable[i] = radiusTable[2 * maxPos - i]
		} else {
			radiusTable[i] = 1
		}
		for ; i + radiusTable[i] < strBufferLen && i - radiusTable[i] >= 0 &&
			tmpStr[i + radiusTable[i]] == tmpStr[i - radiusTable[i]]; {
			radiusTable[i]++
		}
		if radiusTable[i] - 1 > maxRight - maxPos {
			maxRight = i + radiusTable[i] - 1
			maxPos = i
		}
	}
	return s[(maxPos - (radiusTable[maxPos] - 1)) / 2 : (maxPos + radiusTable[maxPos] - 1) / 2]
}

func main() {
	fmt.Print(longestPalindrome("abdedcdf"))
}
