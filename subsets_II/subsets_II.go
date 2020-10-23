package subsets_II

import (
	"fmt"
	"sort"
)

// 具体来说，就是给你一个数组nums，nums.length=n,每次从nums中取出数目不超过n的元素，求所有的取法，也就是求组合

var results [][]int
var t []int

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	results = [][]int{}
	results = append(results, []int{})
	backtracking(0, nums)
	return results
}

func backtracking(cur int, nums []int) {
	if cur == len(nums) {
		return
	}

	for j := cur; j < len(nums); j += 1 {
		if j > cur && nums[j] == nums[j - 1] {
			continue
		}
		t = append(t, nums[j])
		results = append(results, append([]int{}, t...))
		backtracking(j + 1, nums)
		t = t[0:len(t) - 1]
	}
}

func main() {
	fmt.Println(subsetsWithDup([]int{1,1}))
}