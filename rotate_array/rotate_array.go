package rotate_array

func rotate(nums []int, k int)  {
	for start, count := 0, 0; count < len(nums); start += 1 {
		for prev, temp := start, nums[start];;{
			prev = (prev + k) % len(nums)
			nums[prev], temp = temp, nums[prev]
			count += 1
			if prev == start {break}
		}
	}
}



