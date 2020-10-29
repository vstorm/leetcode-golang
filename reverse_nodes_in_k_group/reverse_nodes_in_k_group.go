package reverse_nodes_in_k_group

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h := &ListNode{}
	h.Next = head

	pre, cur, i := h, h.Next, 1

	for ; cur.Next != nil; {
		if i == k {
			pre, cur, i = cur, cur.Next, 1
			continue
		}

		pre.Next, cur.Next, cur.Next.Next = cur.Next, cur.Next.Next, pre.Next
		i += 1

	}

	if i < k {
		cur = pre.Next
		for ; cur.Next != nil; {
			pre.Next, cur.Next, cur.Next.Next = cur.Next, cur.Next.Next, pre.Next
		}
	}

	return h.Next
}
