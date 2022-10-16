package LinkedList

// 反转链表前N个元素

func reverseListN(head *ListNode, n int) *ListNode {
	return reverseN(head, n)
}

var successor *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		// 记录第N+1个节点
		successor = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor // 因为只反转前N个元素, 反转完的节点要指向未反转的部分
	return last
}
