package jump_game_II

func jump(nums []int) int {
	// 我们维护当前能够到达的最大下标位置，记为边界。我们从左到右遍历数组，到达边界时，更新边界并将跳跃次数增加 1
	step, end, reach := 0, 0, 0
	for i := 0; i < len(nums) - 1; i += 1 {
		reach = max(reach, nums[i]+i)
		// 提前结束
		if reach >= len(nums)-1 {
			step += 1
			break
		}
		if i == end {
			step += 1
			end = reach
		}
	}
	return step
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}