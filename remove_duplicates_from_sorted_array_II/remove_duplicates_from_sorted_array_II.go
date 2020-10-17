package remove_duplicates_from_sorted_array_II

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	i, j, count := 1, 1, 1
	for ; j < len(nums); j += 1 {
		if nums[j] == nums[j - 1] {
			count += 1
		} else {
			count = 1
		}

		if count <= 2 {
			nums[i] = nums[j]
			i += 1
		}
	}
	return i
}
