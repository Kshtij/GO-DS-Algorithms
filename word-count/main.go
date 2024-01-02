package main

import (
	"log"
	"sort"
)

func main() {
	log.Println(countFrequencies([]string{"the", "dog", "got", "the", "bone"}))
}

func countFrequencies(words []string) []int {

	wordCountMap := make(map[string]int)
	uniqueWordList := make([]string, 0, 1)
	for i := 0; i < len(words); i++ {
		count, ok := wordCountMap[words[i]]
		if !ok {
			uniqueWordList = append(uniqueWordList, words[i])
		}
		wordCountMap[words[i]] = count + 1
	}

	sort.Strings(uniqueWordList)

	countArr := make([]int, len(uniqueWordList))

	for i := 0; i < len(uniqueWordList); i++ {
		countArr[i] = wordCountMap[uniqueWordList[i]]
	}

	return countArr
}
