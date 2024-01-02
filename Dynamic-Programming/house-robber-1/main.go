package main

func main() {

}

func rob(nums []int) int {
	table := make([]int, len(nums))
	length := len(nums)

	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	table[0] = nums[0]
	table[1] = nums[1]

	if length >= 3 {
		table[2] = nums[2] + nums[0]

		for i := 3; i < len(nums); i++ {
			max := -1
			max = table[i-2]
			if max < table[i-3] {
				max = table[i-3]
			}

			table[i] = nums[i] + max
		}
	}

	maxRevenue := table[length-1]
	if maxRevenue < table[length-2] {
		maxRevenue = table[length-2]
	}

	return maxRevenue
}
