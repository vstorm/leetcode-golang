package increasing_triplet_subsequence

import "math"

func increasingTriplet(nums []int) bool {
	small, mid := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num < small {
			small = num
			continue
		}
		if num > small && num < mid {
			mid = num
			continue
		}
		if num > mid {
			return true
		}
	}
	return false
}
