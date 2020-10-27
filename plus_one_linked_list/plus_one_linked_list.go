package plus_one_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func plusOne(head *ListNode) *ListNode {
	carryF := helper(head)
	if carryF == 1 {
		h := &ListNode{
			Val:  1,
			Next: head,
		}
		return h
	}
	return head
}

func helper(head *ListNode) int {
	if head == nil {
		return 1
	}
	head.Val += helper(head.Next)
	if head.Val == 10 {
		head.Val = 0
		return 1
	}
	return 0
}