package main

var results [][]int
var t []int

func combinationSum(candidates []int, target int) [][]int {
	results = [][]int{}
	backtracking(0, 0, target, candidates)
	return results
}

func backtracking(i, sum, target int, candidates []int) {
	if i == len(candidates) {
		return
	}
	if sum == target {
		results = append(results, append([]int{}, t...))
		return
	}
	if sum > target {
		return
	}

	for j := i; j < len(candidates); j += 1 {
		t = append(t, candidates[j])
		sum += candidates[j]
		backtracking(j, sum, target, candidates)
		t = t[:len(t) - 1]
		sum -= candidates[j]
	}
}

func main() {

}
