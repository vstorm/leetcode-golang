package trapping_rain_water

func trap(height []int) int {
	sum := 0
	l, r := 0, len(height)-1
	lMax, rMax := 0, 0
	for ; l <= r; {		// l <= r 每个位置都要遍历一次
		if lMax < rMax {		// 如果lMax < rMax, 那么左指针l位置的积水量就可以计算了
			if height[l] <= lMax {
				sum += lMax - height[l]
			} else {	// 如果当前位置的高度>lMax, 那么就无法积水了，此时更新lMax
				lMax = height[l]
			}
			l += 1
		} else {	// 同理
			if height[r] <= rMax {
				sum += rMax - height[r]
			} else {
				rMax = height[r]
			}
			r -= 1
		}
	}
	return sum
}