package insert_interval

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	merge := make([][]int, 0)

	i := 0
	for ; i < n &&  intervals[i][1] < newInterval[0]; i += 1 {
		merge = append(merge, intervals[i])
	}

	for end := newInterval[1];i < n && intervals[i][0] <= end; i += 1 {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
	}
	merge = append(merge, newInterval)

	for ; i < n; i += 1 {
		merge = append(merge, intervals[i])
	}
	return merge
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