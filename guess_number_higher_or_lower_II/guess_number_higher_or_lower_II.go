package guess_number_higher_or_lower_II

import "math"

func getMoneyAmount(n int) int {
	results := make([][]int, n+1)

	for i := 0; i <= n; i += 1 {
		for j := 0; j <= n; j += 1 {
			results[i][j] = math.MaxInt32
		}
	}

	for i := 0; i <= n; i += 1 {
		results[i][i] = 0
	}


	for j := 2; j <= n; j += 1 {
		for i := j - 1; i >= 1 ; i -= 1 {
			for k := i+1; k <= j-1; k += 1 {
				results[i][j]=min(k+max(results[i][k-1],results[k+1][j]),results[i][j]);
			}
			results[i][j]=min(results[i][j],i+results[i+1][j]);
			results[i][j]=min(results[i][j],j+results[i][j-1]);
		}
	}
	return results[1][n]
}


func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}