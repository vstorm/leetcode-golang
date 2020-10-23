package combinations

var results [][]int
var t []int

func combine(n int, k int) [][]int {
	results = [][]int{}
	backtracking(0, n, k)
	return results
}

func backtracking(i, n, k int) {
	if len(t) == k {
		results = append(results, append([]int{}, t...))
		return
	}
	for j := i; j < n; j += 1 {
		t = append(t, j + 1)
		backtracking(j + 1, n, k)
		t = t[:len(t) - 1]
	}
}