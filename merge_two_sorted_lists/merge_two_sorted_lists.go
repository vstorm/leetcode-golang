package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur, i1, i2 := head, l1, l2
	for ; i1 != nil && i2 != nil; {
		if i1.Val <= i2.Val {
			cur.Next = i1
			i1 = i1.Next
		} else {
			cur.Next = i2
			i2 = i2.Next
		}
		cur = cur.Next
	}

	if i1 == nil {
		cur.Next = i2
	}
	if i2 == nil {
		cur.Next = i1
	}

	return head.Next
}
