package partition_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	beforeH, afterH := &ListNode{}, &ListNode{}
	bh, ah := beforeH, afterH
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			bh.Next = cur
			bh = bh.Next
		} else {
			ah.Next = cur
			ah = ah.Next
		}
	}

	ah.Next = nil

	bh.Next = afterH.Next
	return beforeH.Next
}