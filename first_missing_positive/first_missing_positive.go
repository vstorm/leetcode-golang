package first_missing_positive

import "math"

func FirstMissingPositive(nums []int) int {
	for i := range nums {
		if nums[i] <= 0 {
			nums[i] = len(nums) + 1
		}
	}
	for _, num := range nums {
		num = int(math.Abs(float64(num)))
		if num < len(nums) {
			nums[num - 1] = -int(math.Abs(float64(nums[num - 1])))
		}
	}
	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}