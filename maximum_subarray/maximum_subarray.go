package maximum_subarray

func maxSubArray(nums []int) int {
	max, pre := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if pre + nums[i] > nums[i] {
			pre = pre + nums[i]
		} else {
			pre = nums[i]
		}
		if pre > max {
			max = pre
		}
	}
	return max
}
