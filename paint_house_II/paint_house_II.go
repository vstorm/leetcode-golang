package paint_house_II

import "math"

func minCostII(costs [][]int) int {
	n := len(costs)
	if n <= 0 {
		return 0
	}

	k := len(costs[0])

	for i, j := n-2, n-1; i >= 0; i, j = i-1, j-1 {
		for l := 0; l < k; l += 1 {
			temp := costs[j][l]
			costs[j][l] = math.MaxInt32
			costs[i][l] += min(costs[j]...)
			costs[j][l] = temp
		}
	}
	return min(costs[0]...)
}

func min(nums ...int) int {
	m := nums[0]
	for _, num := range nums {
		if num < m {
			m = num
		}
	}
	return m
}

