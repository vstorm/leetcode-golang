package palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	i, j := head, head
	for ;j.Next != nil && j.Next.Next != nil; i, j = i.Next, j.Next.Next{}
	for cur := i.Next; cur.Next != nil; cur = cur.Next {
		i.Next, cur.Next, cur.Next.Next = cur.Next, cur.Next.Next, i.Next
	}

	for x, y := head, i.Next; y != nil; x, y = x.Next, y.Next {
		if x.Val != y.Val {
			return false
		}
	}
	return true
}
