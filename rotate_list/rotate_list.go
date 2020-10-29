package rotate_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	h := &ListNode{}
	h.Next = head

	p, q := h, h

	for i := 0; i < k; i += 1 {
		if q.Next == nil {
			q = h.Next
		} else {
			q = q.Next
		}
	}

	for ; q.Next != nil; {
		p, q = p.Next, q.Next
	}

	if p == h || p.Next == nil {
		return h.Next
	}

	h.Next, p.Next, q.Next = p.Next, q.Next, h.Next

	return h.Next
}
