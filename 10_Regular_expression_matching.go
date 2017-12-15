package main

import "fmt"

func isMatch(s string, p string) bool {
	pFoundIndex := 0
	pLength := len(p)
	sLength := len(s)
	sFound := 0
	for i := 0; i < sLength; i++ {
		if pFoundIndex <= pLength-1 {
			// .
			if p[pFoundIndex] == '.' {
				// no overflow
				if pFoundIndex+1 <= pLength-1 && p[pFoundIndex+1] == '*' {
					// .*
					if pFoundIndex+2 > pLength-1 {
						return true
					} else {
						pFoundIndex += 2
						i--
						continue
					}
				}
				sFound++
				pFoundIndex++
				continue
				// char
			} else {
				if pFoundIndex+1 <= pLength-1 && p[pFoundIndex+1] == '*' {
					// is match
					if p[pFoundIndex] == s[i] {
						sFound++
						for i += 1; i < sLength && p[pFoundIndex] == s[i] ; i++ {
							sFound++
						}
					}
					pFoundIndex += 2
					i--
					continue
				} else {
					// found
					if p[pFoundIndex] == s[i] {
						sFound++
						pFoundIndex++
						continue
					}
				}
			}
		}
	}
	if sFound == sLength && pFoundIndex > pLength - 1 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isMatch("aa","a"))
	fmt.Println(isMatch("aa","aa"))
	fmt.Println(isMatch("aaa","aa"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("aa", ".*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("aaa", "ab*a*c*a"))
	fmt.Println(isMatch("aaa", "aaaa"))
}
