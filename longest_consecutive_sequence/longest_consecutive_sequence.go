package longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}

	maxL := 0
	for _, num := range nums {
		if m[num-1] {
			continue
		}
		curL := 0
		for start := num; m[start]; curL, start = curL+1, start+1 {}
		maxL = max(maxL, curL)
	}
	return maxL
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}