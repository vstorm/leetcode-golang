package jump_game

func canJump(nums []int) bool {
	reach, n := 0, len(nums)
	for i := 0; i <= reach && reach < n-1; i += 1 {
		reach = max(reach, nums[i]+i)
	}
	return reach >= n - 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
