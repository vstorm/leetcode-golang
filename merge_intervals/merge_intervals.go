package merge_intervals

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	quickSort(intervals, 0, n-1)
	merge := make([][]int, 0)
	merge = append(merge, intervals[0])
	for _, interval := range intervals {
		m := merge[len(merge)-1]
		if interval[0] <= m[1] {
			if interval[1] > m[1] {
				m[1] = interval[1]
			}
		} else {
			merge = append(merge, interval)
		}
	}
	return merge
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