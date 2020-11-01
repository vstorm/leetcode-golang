package burst_balloons

var cache [][]int

func maxCoins(nums []int) int {
	n := len(nums)
	value := make([]int, n + 2)
	value[0], value[n + 1] = 1, 1
	for i := 1; i <= n; i += 1 {
		value[i] = nums[i - 1]
	}
	cache = make([][]int, n + 2)
	for i := 0; i <= n+1; i += 1 {
		for j := 0; j <= n+1; j += 1 {
			cache[i] = append(cache[i], -1)
		}
	}
	return solve(value, 0, n + 1)
}

func solve(value []int, left, right int) int {
	if left >= right - 1 {
		return 0
	}
	if cache[left][right] != -1 {
		return cache[left][right]
	}
	result := 0
	for i := left + 1; i < right; i += 1 {
		r := value[left] * value[i] * value[right] + solve(value, left, i) + solve(value, i, right)
		result = max(result, r)
	}
	cache[left][right] = result
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}