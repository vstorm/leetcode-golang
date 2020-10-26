package remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	h := &ListNode{}
	h.Next = head

	p, q := h, h

	for i := 0; i < n + 1; i += 1 {
		q = q.Next
	}
	for ; q != nil; p, q = p.Next, q.Next {}
	p.Next = p.Next.Next
	return h.Next
}
