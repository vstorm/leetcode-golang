package meeting_rooms

func canAttendMeetings(intervals [][]int) bool {
	n := len(intervals)
	quickSort(intervals, 0, n-1)
	for i := 0; i < n-1; i += 1 {
		if intervals[i+1][0] < intervals[i][1] {
			return false
		}
	}
	return true
}

// 快速排序
func quickSort(intervals [][]int, left, right int) {
	if left >= right {
		return
	}

	pivot := left+1
	index := partition(intervals, left, right, pivot)	// partition将pivot的位置安排好了

	// 递归使用快速排序，排序pivot左边的元素和右边的元素
	quickSort(intervals, left, index-1)
	quickSort(intervals, index+1,right)
}

func partition(intervals [][]int, left, right, pivot int) int {
	pivotV := intervals[pivot]
	// 先将pivot移到结尾
	intervals[pivot], intervals[right] = intervals[right], intervals[pivot]
	storeI := left
	for i := left; i < right; i += 1 {
		if intervals[i][0] <= pivotV[0] {
			intervals[i], intervals[storeI] = intervals[storeI], intervals[i]
			storeI += 1
		}
	}
	// 将pivot移到storeI的位置，左边的数<=pivot,右边的数>pivot

	intervals[storeI], intervals[right] = intervals[right], intervals[storeI]
	return storeI
}
