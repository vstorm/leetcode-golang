package find_the_duplicate_number

// 41-缺失的第一个正数
func findDuplicate(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		if m[num] {
			return num
		} else {
			m[num] = true
		}
	}
	return 0
}
