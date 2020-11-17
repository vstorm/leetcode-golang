package maximum_gap

import (
	"math"
	"strconv"
)

func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	// 基数排序
	extras := make([][]int, 10)

	maxNum := math.MinInt32
	for _, num := range nums {
		maxNum = max(maxNum, num)
	}
	d := len(strconv.Itoa(maxNum))

	for i := 0; i < d ; i += 1 {
		for _, num := range nums {
			index := (num / int(math.Pow(float64(10), float64(i)))) % 10
			extras[index] = append(extras[index], num)
		}

		j := 0
		for _, e := range extras {
			for _, num := range e {
				nums[j] = num
				j += 1
			}
		}
		extras = make([][]int, 10)
	}

	maxGap := 0
	for i := 0; i < n-1; i += 1 {
		maxGap = max(maxGap, nums[i+1]-nums[i])
	}

	return maxGap
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
