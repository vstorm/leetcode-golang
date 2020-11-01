package paint_house

func minCost(costs [][]int) int {
	n := len(costs)
	if n <= 0 {
		return 0
	}
	for i, j := n-2, n-1; i >= 0; i, j = i-1, j-1 {
		costs[i][0] += min(costs[j][1], costs[j][2])
		costs[i][1] += min(costs[j][0], costs[j][2])
		costs[i][2] += min(costs[j][0], costs[j][1])
	}
	return min(costs[0][0], costs[0][1], costs[0][2])
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
