package main

import "log"

func main() {
	nums := []int{3, 2, 1, 0, 4}
	log.Println("Reachable :", canJump(nums))
}

func canJump(nums []int) bool {
	length := len(nums)
	lastStepToReach := length - 1

	for i := length - 2; i >= 0; i-- {

		if nums[i]+i >= lastStepToReach {
			lastStepToReach = i
		}
	}

	return lastStepToReach == 0
}

/*
	if arr[i] >= lastStepToreach {
		lastStepToReach = i
		i = i-1
	}

*/
