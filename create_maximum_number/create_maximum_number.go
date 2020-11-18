package create_maximum_number

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	m, n := len(nums1), len(nums2)
	maxR := make([]int, k)
	for i := max(0, k-n); i <= min(m, k); i += 1 {
		maxR = maxInts(maxR, merge(pickMax(nums1, m-i), pickMax(nums2, n-(k-i))))
	}
	return maxR
}

func pickMax(nums []int, k int) []int {
	stack := make([]int, 0, len(nums))
	n := k
	for _, num := range nums {
		for ; len(stack) > 0 && k > 0 && stack[len(stack)-1] < num; k -= 1 {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}
	return stack[:len(nums)-n]
}

func merge(nums1 []int, nums2 []int) []int {
	results := make([]int, 0)
	for ;len(nums1) > 0 && len(nums2) > 0; {
		if larger(nums1, nums2) {
			results = append(results, nums1[0])
			nums1 = nums1[1:len(nums1)]
		}  else {
			results = append(results, nums2[0])
			nums2 = nums2[1:len(nums2)]
		}
	}
	if len(nums1) == 0 {
		results = append(results, nums2...)
	} else {
		results = append(results, nums1...)
	}
	return results
}

func larger(nums1, nums2 []int) bool{
	k := min(len(nums1), len(nums2))
	for i := 0; i < k; i += 1 {
		if nums1[i] > nums2[i] {
			return true
		} else if nums1[i] < nums2[i] {
			return false
		}
	}
	if len(nums1) == k {
		return false
	}
	return true
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxInts(nums1, nums2 []int) []int {
	k := min(len(nums1), len(nums2))
	for i := 0; i < k; i += 1 {
		if nums1[i] > nums2[i] {
			return nums1
		} else if nums1[i] < nums2[i] {
			return nums2
		}
	}
	return nums1
}