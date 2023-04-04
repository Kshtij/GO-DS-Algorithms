package main

import "log"

func main() {
	log.Println("Min Distance :", jump([]int{2, 3, 1, 1, 4}))
}

func jump(nums []int) int {
	minDistance := 0
	length := len(nums)
	distance := make([]int, length)
	distance[length-1] = 0
	for i := length - 2; i >= 0; i-- {
		distance[i] = length
		for j := 0; j <= nums[i] && i+j < length; j++ {

			minDistance = distance[i+j] + 1
			if minDistance < distance[i] {
				distance[i] = minDistance
			}

		}
	}
	return distance[0]
}
