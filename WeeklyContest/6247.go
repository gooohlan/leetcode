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

func removeNodes2(head *ListNode) *ListNode {
	var (
		max      int
		traverse func(*ListNode) *ListNode
	)

	traverse = func(node *ListNode) *ListNode {
		if node == nil {
			return nil
		}

		node.Next = traverse(node.Next)
		if max > node.Val {
			return node.Next
		} else {
			max = node.Val
			return node
		}
	}

	head = traverse(head)
	return head
}
