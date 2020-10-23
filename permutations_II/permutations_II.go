package permutations_II

import (
	"sort"
)

var results [][]int
var used []int

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	results = [][]int{}
	used = make([]int, len(nums))
	backtracking(0, nums)
	return results
}

func backtracking(i int, nums []int) {
	if i == len(nums) {
		results = append(results, append([]int{}, nums...))
		return
	}
	m := make(map[int]bool)
	for j := i; j < len(nums); j += 1 {
		if m[nums[j]] {
			continue
		}
		m[nums[j]] = true
		nums[j], nums[i] = nums[i], nums[j]
		backtracking(i + 1, nums)
		nums[j], nums[i] = nums[i], nums[j]
	}
}
