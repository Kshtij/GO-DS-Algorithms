package main

func main() {

}

func maxSlidingWindow(nums []int, k int) []int {
	length := len(nums)
	table := make([][]int, length)

	for i := 0; i < length; i++ {
		table[i] = make([]int, k)
		table[i][0] = nums[i]
	}

	for leap := 2; leap <= k; leap++ {
		for start := 0; start < length-leap+1; start++ {
			elementToConsider := nums[start+leap-1]
			if elementToConsider < table[start][leap-2] {
				elementToConsider = table[start][leap-2]
			}
			table[start][leap-1] = elementToConsider
		}
	}
	list := make([]int, 0, 1)
	for j := 0; j < length-k+1; j++ {
		list = append(list, table[j][k-1])
	}
	return list
}
