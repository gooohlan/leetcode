package LinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

// 定义两个指针,prev与curr, prev为空,curr指向头结点
// 将curr指向prev完成反转, prev和curr依次向后移动
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev // 反转指针
		prev = curr      // 向后移动
		curr = next      // 向后移动

		// curr.Next, prev, curr = prev, curr, curr.Next
	}
	return prev
}

func reverseList2(node *ListNode) *ListNode {
	return reverse(node)
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
