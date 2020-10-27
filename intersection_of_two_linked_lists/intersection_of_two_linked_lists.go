package intersection_of_two_linked_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var tA, tB *ListNode
	pA, pB := headA, headB
	for ;pA != pB; {
		if pA.Next == nil {
			tA = pA
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB.Next == nil {
			tB = pB
			pB = headA
		} else {
			pB = pB.Next
		}

		if tA != nil && tB != nil && tA != tB {
			return nil
		}
	}
	return pA
}
