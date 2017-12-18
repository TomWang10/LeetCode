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

func isMatch2(s string, p string) bool {
	sLength := len(s)
	pLength := len(p)
	if pLength == 0 {
		return sLength == 0
	}
	firstMatch := sLength != 0 && (s[0] == p[0] || p[0] == '.')

	if pLength >= 2 && p[1] == '*' {
		return isMatch2(s, p[2:]) || (firstMatch && isMatch2(s[1:], p))
	} else {
		return firstMatch && isMatch2(s[1:], p[1:])
	}
}

func main() {
	fmt.Println(isMatch2("aa","a"))
	fmt.Println(isMatch2("aa","aa"))
	fmt.Println(isMatch2("aaa","aa"))
	fmt.Println(isMatch2("aa", "a*"))
	fmt.Println(isMatch2("aa", ".*"))
	fmt.Println(isMatch2("ab", ".*"))
	fmt.Println(isMatch2("aab", "c*a*b"))
	fmt.Println(isMatch2("aaa", "ab*a*c*a"))
	fmt.Println(isMatch2("aaa", "aaaa"))
}
