package sort_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 中心点
	i, j := head, head
	for ; j.Next != nil && j.Next.Next != nil; i, j = i.Next, j.Next.Next {
	}

	temp := i.Next
	i.Next = nil

	left := sortList(head)
	right := sortList(temp)

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
