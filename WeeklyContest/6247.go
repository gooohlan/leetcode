package WeeklyContest

type ListNode struct {
	Val  int
	Next *ListNode
}

// 将链表节点放入栈中,遇到更大的则出栈,
func removeNodes(head *ListNode) *ListNode {
	stack := []int{}

	for head != nil {
		for len(stack) > 0 && head.Val > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, head.Val)
		head = head.Next
	}

	head = &ListNode{}
	node := head
	for _, v := range stack {
		head.Next = &ListNode{Val: v}
		head = head.Next
	}
	return node.Next
}
