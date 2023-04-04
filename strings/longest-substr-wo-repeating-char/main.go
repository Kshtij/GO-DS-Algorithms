package main

import "log"

func main() {
	log.Println(lengthOfLongestSubstring("aab"))
}

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	start := 0
	finish := 0
	maxLength := 1
	charSet := make(map[byte]int)
	charSet[s[0]] = 0
	noRepeat := true

	for i := 1; i < length; i++ {
		index, ok := charSet[s[i]]
		if !ok {
			finish++
		} else {
			noRepeat = false
			if finish-start+1 > maxLength {
				maxLength = finish - start + 1
			}
			start = index + 1
			finish++
		}
		charSet[s[i]] = i
	}
	if noRepeat {
		if finish-start+1 > maxLength {
			maxLength = finish - start + 1
		}
	}
	return maxLength
}
