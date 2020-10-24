package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// 哈希表
//func hasCycle(head *ListNode) bool {
//	m := make(map[*ListNode]bool)
//	for cur := head; cur != nil; cur = cur.Next {
//		if m[cur] {
//			return true
//		}
//		m[cur] = true
//	}
//	return false
//}

// 快慢指针（快指针每次走两步，慢指针每次走一步，不管单数环还是双树环，快慢指针总会在某一时刻相遇)
func hasCycle(head *ListNode) bool {
	if head == nil && head.Next == nil {
		return false
	}
	for fast, slow := head, head;;{
		if fast.Next == nil || fast.Next.Next == nil {
			break
		}
		slow = slow.Next
		if slow == nil {
			break
		}
		if fast == slow {
			return true
		}
	}
	return false
}