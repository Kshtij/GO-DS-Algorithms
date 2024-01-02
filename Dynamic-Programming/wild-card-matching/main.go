package main

import (
	"log"
	"time"
)

func main() {
	t := time.Now()
	//s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	//p := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	s := "adceb"
	p := "*a*b"
	//log.Println(isMatch(s, p))
	log.Println(bottomUp(s, p))
	log.Println("time taken :", time.Now().Sub(t))
}

// func isMatch(s string, p string) bool {
// 	memo := make(map[string]bool)
// 	return matchHelper(s, p, memo)
// }

func matchHelper(s string, p string, memo map[string]bool) bool {
	var retVal bool

	retVal, ok := memo[s+"-"+p]
	if ok {
		return retVal
	}

	if s == "" && p == "" {
		retVal = true
	} else if s != "" && p == "" {
		retVal = false
	} else if s == "" {
		if p[0] == '*' {
			retVal = matchHelper(s, slice(p), memo)
		} else {
			retVal = false
		}
	} else if p[0] == s[0] {
		retVal = matchHelper(slice(s), slice(p), memo)
	} else if p[0] == '?' {
		retVal = matchHelper(slice(s), slice(p), memo)
	} else if p[0] == '*' {
		if s == "" {
			retVal = matchHelper(s, slice(p), memo)
		} else {
			retVal = matchHelper(slice(s), p, memo) || matchHelper(slice(s), slice(p), memo) || matchHelper(s, slice(p), memo)
		}

	}

	log.Println("entering in map :", s+"-"+p)
	memo[s+"-"+p] = retVal
	return retVal
}

func slice(s string) string {
	if len(s) <= 1 {
		return ""
	}
	return s[1:]
}

func isMatch(s string, p string) bool {
	return bottomUp(s, p)
}

func bottomUp(s, p string) bool {
	matrix := make([][]bool, len(p)+1)
	for i := 0; i < len(p)+1; i++ {
		matrix[i] = make([]bool, len(s)+1)
	}
	matrix[0][0] = true

	for i := 1; i <= len(s); i++ {
		matrix[0][i] = false
	}

	for i := 1; i <= len(p); i++ {
		if p[i-1] == '*' {
			matrix[i][0] = true && matrix[i-1][0]
		} else {
			matrix[i][0] = false
		}
	}

	for i := 1; i <= len(p); i++ {
		for j := 1; j <= len(s); j++ {
			if s[j-1] == p[i-1] {
				matrix[i][j] = true && matrix[i-1][j-1]
			} else if p[i-1] == '*' {
				matrix[i][j] = true && (matrix[i-1][j-1] || matrix[i][j-1] || matrix[i-1][j])
			} else if p[i-1] == '.' && matrix[i-1][j-1] == true {
				matrix[i][j] = true
			}
		}
	}

	return matrix[len(p)][len(s)]

}
