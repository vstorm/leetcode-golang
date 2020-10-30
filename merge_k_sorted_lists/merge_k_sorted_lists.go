package merge_k_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}
	return merge(lists, 0, l-1)
}

func merge(lists []*ListNode, lo, hi int) *ListNode {
	if lo == hi {
		return lists[lo]
	}
	mid := (lo + hi) / 2
	left := merge(lists, lo, mid)
	right := merge(lists, mid+1, hi)

	h := &ListNode{}
	cur := h
	for ; left != nil && right != nil; cur = cur.Next {
		if left.Val <= right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
	}

	if left == nil {
		cur.Next = right
	} else {
		cur.Next = left
	}

	return h.Next
}
