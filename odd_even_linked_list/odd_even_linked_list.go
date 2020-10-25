package odd_even_linked_list


type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h := &ListNode{}
	h.Next = head
	for odd, even := h.Next, h.Next.Next;; {
		if even == nil || even.Next == nil {
			break
		}
		odd.Next, even.Next, even.Next.Next = even.Next, even.Next.Next, odd.Next
		odd, even = odd.Next, even.Next
	}
	return h.Next
}
