package reorder_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil {
		return
	}

	i, j := head, head
	for ; j.Next != nil && j.Next.Next != nil; i, j = i.Next, j.Next.Next {}

	for cur := i.Next; cur.Next != nil; {
		i.Next, cur.Next, cur.Next.Next = cur.Next, cur.Next.Next, i.Next
	}

	for x, y := head, i; x != y;{
		x, x.Next, y.Next, y.Next.Next = x.Next, y.Next, y.Next.Next, x.Next
	}
}