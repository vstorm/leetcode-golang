package reverse_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode

	for cur := head; cur != nil; {
		cur.Next, cur, pre = pre, cur.Next, cur

	}
	return pre
}