package edit_distance

func minDistance(word1 string, word2 string) int {
	m := len(word1) + 1
	n := len(word2) + 1

	result := make([][]int, 0)
	for i := 0; i < n; i += 1 {
		l := make([]int, m)
		result = append(result, l)
	}
	for i := 0; i < m; i += 1 {
		result[0][i] = i
	}
	for i := 0; i < n; i += 1 {
		result[i][0] = i
	}




	for i := 1; i < n; i += 1 {
		for j := 1; j < m; j += 1 {
			var t int
			if word1[j-1] == word2[i-1] {
				t = result[i-1][j-1]
			} else {
				t = result[i-1][j-1] + 1
			}
			result[i][j] = min(result[i][j-1]+1, result[i-1][j]+1, t)
		}
	}
	return result[n-1][m-1]
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
}}