package minimum_path_sum


func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	results := make([]int, n)
	var r int

	for i := 0; i < m; i += 1 {
		for j := 0; j < n; j += 1 {
			if j == 0 {
				r = grid[i][j] + results[j]
			} else if i == 0 {
				r = grid[i][j] + r
			} else {
				r = min(grid[i][j] + r, grid[i][j] + results[j])
			}
			results[j] = r
		}
	}
	return r
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}