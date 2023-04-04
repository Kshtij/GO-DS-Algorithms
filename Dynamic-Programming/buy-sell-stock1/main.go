package main

func main() {

}

func maxProfit(prices []int) int {
	buy := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if buy > prices[i] {
			buy = prices[i]
		} else if prices[i]-buy > maxProfit {
			maxProfit = prices[i] - buy
		}
	}

	return maxProfit
}
