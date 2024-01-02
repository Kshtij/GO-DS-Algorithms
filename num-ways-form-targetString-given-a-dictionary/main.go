package main

import "log"

func main() {
	words := []string{"cbabddddbc", "addbaacbbd", "cccbacdccd", "cdcaccacac", "dddbacabbd", "bdbdadbccb", "ddadbacddd", "bbccdddadd", "dcabaccbbd", "ddddcddadc", "bdcaaaabdd", "adacdcdcdd", "cbaaadbdbb", "bccbabcbab", "accbdccadd", "dcccaaddbc", "cccccacabd", "acacdbcbbc", "dbbdbaccca", "bdbddbddda", "daabadbacb", "baccdbaada", "ccbabaabcb", "dcaabccbbb", "bcadddaacc", "acddbbdccb", "adbddbadab", "dbbcdcbcdd", "ddbabbadbb", "bccbcbbbab", "dabbbdbbcb", "dacdabadbb", "addcbbabab", "bcbbccadda", "abbcacadac", "ccdadcaada", "bcacdbccdb"}
	target := "bcbbcccc"

	log.Println("NumWays :", numWays(words, target))
}

func numWays(words []string, target string) int {
	if target == "" {
		return 0
	}

	totalWays := 0
	for i := 0; i < len(words[0]); i++ {
		if len(words[0])-i >= len(target) {
			totalWays += numWaysHelper(words, target, i)
		}

	}

	return totalWays

}

func numWaysHelper(words []string, target string, index int) int {
	//log.Println("rcvd :", target, index)
	if target == "" {
		return 1
	}

	if index >= len(words[0]) {
		return 0
	}

	if len(target) > len(words[0][index:]) {
		return 0
	}

	totalCount := 0
	for _, word := range words {
		if word[index] == target[0] {

			if len(target) > 1 {
				for j := index + 1; j < len(word); j++ {

					totalCount += numWaysHelper(words, target[1:], j)

				}
			} else if len(target) == 1 {
				totalCount = totalCount + 1
			}

		}
	}

	//log.Println("returning :", target, index, totalCount)
	return totalCount
}
