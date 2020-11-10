package contains_duplicate_III

func getId(x, w int) int {
	if x < 0 {
		return (x + 1) / w - 1	// x为负数时
	}
	return x / w
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	m := make(map[int]int)
	w := t + 1 // 注意
	for i, num := range nums {
		id := getId(num, w)
		// k个数，如果有两个数组在同一个桶内(一个桶容纳t个元素, 这样一个桶内的任两个数，差的绝对值肯定小于等于 t)
		if _, ok := m[id]; ok {
			return true
		}
		if _, ok := m[id-1]; ok {
			if num - m[id-1] <= t {
				return true
			}
		}
		if _, ok := m[id+1]; ok {
			if m[id+1] - num  <= t {
				return true
			}
		}
		m[id] = num
		if i >= k {
			delete(m, getId(nums[i-k], w))
		}
	}
	return false
}