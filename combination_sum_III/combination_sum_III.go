package main

import "fmt"

var results [][]int
var t []int

func combinationSum3(k int, n int) [][]int {
	results = [][]int{}
	backtracking(0, k, 0, n)
	return results
}

func backtracking(i, k, sum, n int) {
	if len(t) == k {
		if sum == n {
			results = append(results, append([]int{}, t...))
		}
		return
	}
	for j := i; j < 9; j += 1 {
		v := j + 1
		t = append(t, v)
		sum += v
		backtracking(j + 1, k, sum, n)
		t = t[:len(t) - 1]
		sum -= v
	}
}

func main() {
	fmt.Println(combinationSum3(3, 9))
}