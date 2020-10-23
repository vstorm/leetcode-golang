package permutations

var results [][]int
var t []int

func permute(nums []int) [][]int {
	results = [][]int{}
	backtracking(0, nums)
	return results
}

func backtracking(i int, nums []int) {
	if i == len(nums) {
		results = append(results, append([]int{}, t...))
		return
	}
	for j := i; j < len(nums); j += 1 {
		t = append(t, nums[j])
		nums[j], nums[i] = nums[i], nums[j]
		backtracking(i + 1, nums)
		t = t[:len(t) - 1]
		nums[j], nums[i] = nums[i], nums[j]
	}
}
