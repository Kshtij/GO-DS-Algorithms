package main

import "log"

/*

[1,-3,4,5,-8,7,5,6,-3,9,8,-25,100]
*/
func main() {
	arr := []int{-8, 4, 5, -2, -3, -7, 1000, -4000}
	log.Println("Max Sum :", maxSubArray(arr))
}

func maxSubArray(nums []int) int {

	sum := 0
	maxSum := nums[0] // set it to min number in array

	maxElement := maxSum
	current := 0
	for current < len(nums) {
		if nums[current]+sum > 0 {
			sum = sum + nums[current]
			if sum > maxSum {
				maxSum = sum
			}
		} else {
			sum = 0
		}
		if maxElement < nums[current] {
			maxElement = nums[current]
		}

		current++
	}

	if maxSum < maxElement {
		maxSum = maxElement
	}
	return maxSum
}
