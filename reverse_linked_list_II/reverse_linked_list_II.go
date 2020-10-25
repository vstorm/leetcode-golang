package reverse_linked_list_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m-n == 0 {
		return head
	}

	h := &ListNode{}
	h.Next = head

	for i, j, pre, cur := 1, 0, h, h.Next;; i += 1 {
		if i < m {
			pre = pre.Next
			cur = cur.Next
		}

		if i >= m {
			pre.Next, cur.Next, cur.Next.Next = cur.Next, cur.Next.Next, pre.Next
			j += 1
		}
		if j == n-m {
			break
		}
	}
	return h.Next
}