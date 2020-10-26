package remove_linked_list_elements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	h := &ListNode{}
	h.Next = head

	for pre, cur := h, head; cur != nil; cur = cur.Next{
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = pre.Next
		}
	}
	return h.Next
}