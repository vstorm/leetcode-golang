package median_of_two_sorted_arrays

import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	l := m + n + 1

	result := 0.0
	for start, end := 0, m; start <= end; {
		i := (start + end) / 2
		j := l/2 - i

		var nums1L int
		if i == 0 {
			nums1L = math.MinInt32
		} else {
			nums1L = nums1[i-1]
		}
		var nums1R int
		if i == m {
			nums1R = math.MaxInt32
		} else {
			nums1R = nums1[i]
		}
		var nums2L int
		if j == 0 {
			nums2L = math.MinInt32
		} else {
			nums2L = nums2[j-1]
		}
		var nums2R int
		if j == n {
			nums2R = math.MaxInt32
		} else {
			nums2R = nums2[j]
		}

		if nums1L > nums2R {
			end = i - 1
			continue
		}
		if nums2L > nums1R {
			start = i + 1
			continue
		}

		// 偶数个
		if (m+n) % 2 == 0 {
			result = (math.Max(float64(nums1L), float64(nums2L)) + math.Min(float64(nums1R), float64(nums2R))) / 2
		} else {
			result = math.Max(float64(nums1L), float64(nums2L))
		}
		break
	}
	return result
}

