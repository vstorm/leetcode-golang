package h_index

//func hIndex(citations []int) int {
//	sort.Ints(citations)
//	n := len(citations)
//	h := 0
//	for i := 0; i < n; i += 1 {
//		if n - i <= citations[i] {
//			h = n - i
//			break
//		}
//	}
//	return h
//}

func hIndex(citations []int) int {
	n := len(citations)
	papers := make([]int, n+1)
	for _, c := range citations {
		papers[min(n, c)] += 1
	}
	k := n
	for s := papers[k]; k > s; s += papers[k] {
		k -= 1
	}
	return k
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}