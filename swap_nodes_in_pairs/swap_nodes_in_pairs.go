package swap_nodes_in_pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	h := &ListNode{}
	h.Next = head
	cur := h

	for ;; {
		if cur.Next == nil || cur.Next.Next == nil {
			break
		}
		cur.Next, cur.Next.Next.Next, cur.Next.Next = cur.Next.Next, cur.Next, cur.Next.Next.Next
		cur = cur.Next.Next
	}
	return h.Next
}