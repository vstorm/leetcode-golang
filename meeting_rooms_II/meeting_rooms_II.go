package meeting_rooms_II

import "container/heap"

func minMeetingRooms(intervals [][]int) int {
	n := len(intervals)
	quickSort(intervals, 0, n-1)
	rooms := 0
	var queue Queue
	heap.Init(&queue)
	for i := 0; i < n; i += 1 {
		if len(queue) == 0 {
			heap.Push(&queue, intervals[i][1])
			rooms += 1
			continue
		}
		if queue[0] <= intervals[i][0] {
			heap.Pop(&queue)
		} else {
			rooms += 1
		}
		heap.Push(&queue, intervals[i][1])
	}
	return rooms
}

type Queue []int

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	return q[i] < q[j]
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*q = append(*q, x.(int))
}

func (q *Queue) Pop() interface{} {
	n := len(*q)
	x := (*q)[n-1]
	*q = (*q)[0 : n-1]
	return x
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
