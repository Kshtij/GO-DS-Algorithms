package main

import (
	"log"
	"time"
)

func main() {
	s1 := "aabc"
	s2 := "abad"
	s3 := "aabcabad"
	t := time.Now()
	log.Println(isInterleave(s1, s2, s3))
	log.Println(time.Now().Sub(t))
}

func isInterleave(s1 string, s2 string, s3 string) bool {

}

func helper(s1 string, s2 string, s3 string, matrix [][]bool, s1Index, s2Index, s3Index int) bool {

	for i := s3Index; i < len(s3); i++ {
		if s1Index == len(s1) && s2Index == len(s2) && matrix[s1Index][s2Index] {
			return true
		}
		if s3[s3Index] == s1[s1Index-1] && s3[s3Index] == s2[s2Index-1] {

		} else if s3[s3Index] == s2[s2Index-1] {
			s2Index++
			matrix[s1Index][s2Index] = true
		} else if s3[s3Index] == s1[s1Index-1] {
			s1Index++
			matrix[s1Index][s2Index] = true
		}
	}
}
