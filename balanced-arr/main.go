package main

import "log"

func main() {
	log.Println("IS Balanced :", IsBalanced([]int{3, 5, 2, 4, 3, 2, 1, 2}))
}

func IsBalanced(arr []int) bool {
	//sum := getSum(arr)
	//expectedSum := sum / 2
	selectionSort(arr)
	log.Println("sorted array :", arr)

	firstArr := make([]int, 0, 1)
	secondArr := make([]int, 0, 1)

	start := 0
	end := len(arr) - 1
	flag := true
	for end > start && end-start > 1 {
		log.Println(start, end, flag)
		if flag {
			firstArr = append(firstArr, arr[start])
			firstArr = append(firstArr, arr[end])
		} else {
			secondArr = append(secondArr, arr[start])
			secondArr = append(secondArr, arr[end])
		}

		flag = !flag
		start++
		end--

	}

	if end-start == 1 {
		log.Println(start, end)
		firstArr = append(firstArr, arr[start])
		secondArr = append(secondArr, arr[end])
		start++
		end--
	}

	log.Println("firstArr :", firstArr)
	log.Println("secondArr :", secondArr)

	if getSum(firstArr) == getSum(secondArr) && len(firstArr) == len(secondArr) {
		return true
	}

	return false

}

func getSum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum
}

func selectionSort(arr []int) {

	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
