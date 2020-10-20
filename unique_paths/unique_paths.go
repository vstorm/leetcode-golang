package unique_paths

func uniquePaths(m int, n int) int {
	paths := make([]int, n)
	result := 0
	for i:= 0; i < m; i+= 1 {
		for j:=0; j < n; j+= 1 {
			if i == 0 || j == 0 {
				result = 1
			} else {
				result = result + paths[j]
			}
			paths[j] = result
		}
	}
	return result
}