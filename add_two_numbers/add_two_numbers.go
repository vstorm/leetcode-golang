package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h := l1
	tail := &ListNode{}
	carry := 0
	i, j := l1, l2
	for iv, jv := 0, 0; i != nil || j != nil; {

		if i == nil {
			iv = 0
		} else {
			iv = i.Val
		}
		if j == nil {
			jv = 0
		} else {
			jv = j.Val
		}

		val := iv + jv + carry
		carry = val / 10
		val = val % 10

		if i != nil {
			i.Val, tail, h = val, i, l1
			i = i.Next
		}
		if j != nil {
			j.Val, tail, h = val, j, l2
			j = j.Next
		}
	}

	if carry == 1 {
		tail.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return h
}
