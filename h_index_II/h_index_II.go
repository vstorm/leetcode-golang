package h_index_II

func hIndex(citations []int) int {
	// 因为已经排好序，我们要找的满足n-k>=citations[k](0 <=k<=n-1),最大的k。最后输出n-k
	n := len(citations)
	lo, hi := 0, n-1
	mid := 0
	for ; lo <= hi; {
		mid = (lo + hi) / 2
		if n-mid > citations[mid] {
			lo = mid + 1
		} else if n-mid < citations[mid] {
			hi = mid - 1
		} else {
			return n - mid
		}
	}
	return n - lo
}
