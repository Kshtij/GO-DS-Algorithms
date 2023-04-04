package main

import (
	"log"
	"time"
)

func main() {
	t := time.Now()
	log.Println("Ways :", numDecodings("11011"))
	log.Println("time taken :", time.Now().Sub(t))
}

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	var cache = make(map[string]int)
	cache["0"] = 0
	return CountWays(s, cache)

}

func CountWays(s string, cache map[string]int) int {
	val, ok := cache[s]
	if ok {
		return val
	}

	if len(s) <= 1 {
		return 1
	}
	if s[0] == '0' {
		return 0
	}
	var count int
	if s[len(s)-1] != '0' {
		count = CountWays(s[:len(s)-1], cache)
	}

	if len(s) >= 2 {
		//if s[len(s)-2] == '1' || ((s[len(s)-2] == '2') && (s[len(s)-1] == '0' || s[len(s)-1] == '1' || s[len(s)-1] == '2' || s[len(s)-1] == '3' || s[len(s)-1] == '4' || s[len(s)-1] == '5' || s[len(s)-1] == '6')) {
		if s[len(s)-2] == '1' || ((s[len(s)-2] == '2') && (int(s[len(s)-1]) >= 48 && int(s[len(s)-1]) <= 54)) {
			count = count + CountWays(s[:len(s)-2], cache)
		}
	}
	cache[s] = count
	return count
}
