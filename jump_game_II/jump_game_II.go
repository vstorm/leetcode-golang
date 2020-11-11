package jump_game_II

func jump(nums []int) int {
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