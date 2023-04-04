package main

import "log"

func main() {
	arr := []int{3, 0, 2, 0, 4}
	log.Println("Water Trapped :", trap(arr))
}

func trap(height []int) int {
	firstWall := 0
	for height[firstWall] == 0 && firstWall < len(height) {
		firstWall++
	}

	waterUnit := 0

	for current := firstWall + 2; current < len(height); current++ {
		if height[current] > height[current-1] {
			distance := 1
			maxHeight := height[current-1]
			for wall := current - 2; wall >= firstWall; wall-- { //&& height[wall] < height[current]
				if height[wall] > maxHeight {
					min := getMin(height[current], height[wall])
					waterUnit = waterUnit + min*distance - maxHeight*distance
					if maxHeight < min {
						maxHeight = min
					}
					if height[wall] >= height[current] {
						break
					}
				}

				distance++
			}
		}
	}

	return waterUnit

}

func getMin(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
