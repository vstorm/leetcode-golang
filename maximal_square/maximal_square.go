package maximal_square

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	dp := make([][]int, 0)
	for i := 0; i < m; i += 1 {
		l := make([]int, n)
		dp = append(dp, l)
	}

	r := 0
	for i := 0; i < m; i += 1 {
		for j := 0; j < n; j += 1 {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
				}
				r = max(dp[i][j], r)
			}

		}
	}
	return r * r
}

func min(x, y, z int) int {
	m := x
	if y < m {
		m = y
	}
	if z < m {
		m = z
	}
	return m
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}