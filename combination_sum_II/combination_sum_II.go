package combination_sum_II

import "sort"

var results [][]int
var t []int

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	results = [][]int{}
	backtracking(0, 0, target, candidates)
	return results
}

func backtracking(i, sum, target int, candidates []int) {
	if sum == target {
		results = append(results, append([]int{}, t...))
		return
	}
	if sum > target {
		return
	}
	if i == len(candidates) {
		return
	}
	for j := i; j < len(candidates); j += 1 {
		if j > i && candidates[j] == candidates[j-1] {
			continue
		}
		t = append(t, candidates[j])
		sum += candidates[j]
		backtracking(j+1, sum, target, candidates)
		t = t[:len(t) - 1]
		sum -= candidates[j]
	}
}
