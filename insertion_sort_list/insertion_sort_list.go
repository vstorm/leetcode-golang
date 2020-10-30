package insertion_sort_list

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	h := &ListNode{Val: math.MinInt32}
	h.Next = head

	for pre, cur := h, h.Next; cur != nil; {
		if cur.Val < pre.Val {
			for temp := h; temp.Next != cur ; temp = temp.Next {
				if temp.Next.Val > cur.Val {
					pre.Next, temp.Next, cur.Next = pre.Next.Next, cur, temp.Next
					cur = pre.Next
					break
				}
			}
		} else {
			pre, cur = pre.Next, cur.Next
		}
	}

	return h.Next
}
