package perfect_squares

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	results := make([]int, amount + 1)
	results[0] = 0

	for i := 1; i <= amount; i += 1 {
		minR := math.MaxInt32
		for _, coin := range coins {
			if coin <= i && results[i - coin] != -1 {
				r := results[i - coin] + 1
				minR = min(r, minR)
			}
		}
		if minR == math.MaxInt32 {
			minR = -1
		}
		results[i] = minR
	}
	return results[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}