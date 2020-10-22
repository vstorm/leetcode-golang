package triangle

import "math"

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	results := make([]int, m + 1)

	for i := 0; i < m; i += 1 {
		minR := math.MaxInt32
		for j := len(triangle[i]) - 1; j >= 0; j -= 1 {
			num := triangle[i][j]
			if j == 0 {
				results[j] = results[j] + num
			} else if j == len(triangle[i]) - 1 {
				results[j] = results[j - 1] + num
			} else {
				results[j] = min(results[j - 1] + num, results[j] + num)
			}
			minR = min(minR, results[j])

		}
	}
	return results[m - 1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}