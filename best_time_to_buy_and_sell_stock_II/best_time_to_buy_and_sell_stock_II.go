package best_time_to_buy_and_sell_stock_II

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxP := 0
	for i := 0; i < len(prices); i += 1 {
		if prices[i] > minPrice {
			maxP += prices[i]
		}
		minPrice = prices[i]
	}
	return maxP
}
