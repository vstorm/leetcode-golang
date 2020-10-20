package unique_paths_II

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := 0
	if m > 0 {
		n = len(obstacleGrid[0])
	} else {
		n = 0
	}

	paths := make([]int, n)
	result := 0

	if obstacleGrid[0][0] == 1 {
		paths[0] = 1
	}

	for i := 0; i < m; i += 1 {
		for j := 0; j < n; j += 1 {
			if obstacleGrid[i][j] == 1 {
				result = 0
			} else {
				if j == 0 {
					result = paths[j]
				} else {
					result = result + paths[j]
				}
			}
			paths[j] = result
		}
	}
	return result
}
