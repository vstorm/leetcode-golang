package subsets

var result [][]int
var t []int

func subsets(nums []int) [][]int {
	result = [][]int{}
	helper(0, nums)
	return result
}

func helper(cur int, nums []int) {
	if cur == len(nums) {
		result = append(result, append([]int{}, t...))
		return
	}
	t = append(t, nums[cur])

	cur += 1
	helper(cur, nums)
	t = t[0:len(t) - 1]

	helper(cur, nums)
}