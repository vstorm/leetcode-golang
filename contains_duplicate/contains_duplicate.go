package contains_duplicate

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, num := range nums {
		if m[num] == 1 {
			return true
		} else {
			m[num] = 1
		}
	}
	return false
}
