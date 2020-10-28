package linked_list_cycle_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var meet *ListNode
	for fast, slow := head, head;; {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			meet = fast
			break
		}
	}

	cur := head
	for ; cur != meet; cur, meet = cur.Next, meet.Next {}
	return cur
}