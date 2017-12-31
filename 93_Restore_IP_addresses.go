package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddressesHelper(result *[]string, s string, index int, sub string) {
	if len(s) > (4 - index) * 3 {
		return
	}
	if len(s) < 4 - index {
		return
	}
	// is correct ip
	if len(s) == 0 && index == 4 {
		*result = append(*result, sub[:len(sub)-1])
		return
	}

	for i := 0; i < 3 && i < len(s); i++ {
		if i > 0 && s[0] == '0' {
			return
		}
		digit, _ := strconv.Atoi(s[:i+1])
		if digit > 255 {
			return
		}
		restoreIpAddressesHelper(result, s[i+1:], index+1, sub + s[:i+1] + ".")
	}
}

func restoreIpAddresses(s string) []string {
	var result []string
	restoreIpAddressesHelper(&result, s, 0, "")
	return result
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("010010"))
}
