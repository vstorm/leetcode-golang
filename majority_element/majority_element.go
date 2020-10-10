package majority_element

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num] += 1
		if m[num] > len(nums) / 2 {
			return num
		}
	}
	return 0
}
