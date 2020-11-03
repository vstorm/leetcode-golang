package interleaving_string

func isInterleave(s1 string, s2 string, s3 string) bool {
	m := len(s1)
	n := len(s2)
	if m+n != len(s3) {
		return false
	}
	result := make([][]bool, 0)
	for i := 0; i < n+1; i += 1 {
		l := make([]bool, m+1)
		result = append(result, l)
	}

	result[0][0] = true
	for i := 0; i < n+1; i += 1 {
		for j := 0; j < m+1; j += 1 {
			p := i + j - 1
			if i > 0 && j > 0 {
				result[i][j] = (result[i-1][j] && (s3[p] == s2[i-1])) || (result[i][j-1] && (s3[p] == s1[j-1]))
			} else if i > 0 {
				result[i][j] = result[i-1][j] && (s3[p] == s2[i-1])
			} else if j > 0 {
				result[i][j] = result[i][j-1] && (s3[p] == s1[j-1])
			}
		}
	}
	return result[n][m]
}
