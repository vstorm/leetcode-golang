package container_with_most_water

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxA := (r - l) * min(height[l], height[r])
	for ; l < r; {
		if height[l+1] > height[l] {
			maxA = max(maxA, (r-l-1)*min(height[l+1], height[r]))
			l += 1
			continue
		}
		if height[r-1] > height[r] {
			maxA = max(maxA, (r-1-l)*min(height[l], height[r-1]))
			r -= 1
			continue
		}
		break
	}
	return maxA
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
