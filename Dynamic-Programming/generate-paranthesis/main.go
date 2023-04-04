package main

import "log"

func main() {
	log.Println(generateParenthesis(4))
}

func generateParenthesis(n int) []string {
	openBracket := []string{}
	closeBracket := []string{}

	for i := 0; i < n; i++ {
		openBracket = append(openBracket, "(")
		closeBracket = append(closeBracket, ")")
	}

	return paranthesis(openBracket, closeBracket)
}

func paranthesis(start, end []string) []string {

	if len(start) == 0 && len(end) >= 0 {
		finalStr := ""
		for _, str := range end {
			finalStr = finalStr + str
		}
		return []string{finalStr}
	}

	if len(start) == len(end) {
		prefix := start[len(start)-1]
		suffixArr := paranthesis(start[:len(start)-1], end)

		retArr := make([]string, 0, 1)

		for _, str := range suffixArr {
			newStr := prefix + str
			retArr = append(retArr, newStr)
		}

		return retArr

	} else {

		firstPrefix := start[len(start)-1]
		firstSuffixArr := paranthesis(start[:len(start)-1], end)

		retArr := make([]string, 0, 1)

		for _, str := range firstSuffixArr {
			newStr := firstPrefix + str
			retArr = append(retArr, newStr)
		}

		secondPrefix := end[len(end)-1]

		secondSuffixArr := paranthesis(start, end[:len(end)-1])

		for _, str := range secondSuffixArr {
			newStr := secondPrefix + str
			retArr = append(retArr, newStr)
		}

		return retArr
	}
}
