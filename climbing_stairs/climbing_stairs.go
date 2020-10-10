package climbing_stairs

func climbStairs(n int) int {
	preOne, preTwo, nums := 0, 0, 1
	for i:= 1; i <= n; i++ {
		preTwo = preOne
		preOne = nums
		nums =  preOne + preTwo
	}
	return nums
}