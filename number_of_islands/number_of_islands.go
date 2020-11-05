package number_of_islands

func numIslands(grid [][]byte) int {
	islandNum := 0
	m := len(grid)
	if m == 0 {
		return islandNum
	}
	n := len(grid[0])
	for i := 0; i < m; i += 1 {
		for j := 0; j < n; j += 1 {
			if grid[i][j] == '1' {
				islandNum += 1
				dfs(i, j, grid)
			}
		}
	}
	return islandNum
}

func dfs(i, j int, grid [][]byte) {
	grid[i][j] = '0'

	m := len(grid)
	n := len(grid[0])
	if i > 0 && grid[i-1][j] == '1' {dfs(i-1, j, grid)}
	if j > 0 && grid[i][j-1] == '1' {dfs(i, j-1, grid)}
	if (i < m - 1) && (grid[i+1][j] == '1') {dfs(i+1, j, grid)}
	if (j < n - 1) && (grid[i][j+1] == '1') {dfs(i, j+1, grid)}
}