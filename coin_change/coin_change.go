package coin_change

import "math"

func coinChange(coins []int, amount int) int {
	results := make([]int, amount + 1)
	results[0] = 0

	for i := 1; i < amount; i += 1 {
		minR := math.MaxInt32
		for _, coin := range coins {
			if coin <= i {
				r := results[i - coin] + 1
				minR = min(r, minR)
			}
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