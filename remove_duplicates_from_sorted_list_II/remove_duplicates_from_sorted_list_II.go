package remove_duplicates_from_sorted_list_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	h := &ListNode{}
	h.Next = head
	for i, pre, cur := 0, h, head; cur != nil; {
		for ; cur.Next != nil; {
			if cur.Val == cur.Next.Val {
				cur = cur.Next
				i += 1
			} else {
				break
			}
		}
		if i > 0 {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		i = 0
		cur = cur.Next
	}
	return h.Next
}