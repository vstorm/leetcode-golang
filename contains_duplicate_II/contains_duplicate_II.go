package contains_duplicate_II

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, num := range nums {
		if index, ok := m[num]; ok {
			if i - index <= k {
				return true
			}
		}
		m[num] = i
	}
	return false
}
