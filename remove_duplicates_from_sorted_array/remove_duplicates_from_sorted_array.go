package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	i, j := 0, 1
	for ; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i += 1
			nums[i] = nums[j]
		}
	}
	return i + 1
}
