package LinkedList

func reverseKGroup(head *ListNode, k int) *ListNode {
	tail := head
	for i := 0; i < k; i++ {
		if tail == nil { // 不足K不用返回直接return
			return head
		}
		tail = tail.Next
	}
	// 翻转0,k
	newHead := reverse25(head, tail)
	head.Next = reverseKGroup(tail, k) // 翻转K到2K
	return newHead
}

func reverse25(head, tail *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	curr := head

	for curr != tail {
		curr.Next, prev, curr = prev, curr, curr.Next
	}
	return prev
}
