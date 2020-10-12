package remove_element

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums) - 1
	for ; left < right; {
		if nums[left] == val {
			nums[left] = nums[right]
			right -= 1
		} else {
			left += 1
		}
	}
	return left - 1
}