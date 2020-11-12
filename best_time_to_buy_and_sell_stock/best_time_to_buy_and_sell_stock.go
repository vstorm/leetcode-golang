package best_time_to_buy_and_sell_stock

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxP := 0
	for i := 0; i < len(prices); i += 1 {
		minPrice = min(minPrice, prices[i])
		maxP = max(maxP, prices[i]-minPrice)
	}
	return maxP
}



func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}