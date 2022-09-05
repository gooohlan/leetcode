package linked_list

// 分别取出两个链表的当前值比较,较小的赋值给新链表
// 出事化时需要声明一个虚拟节点,避免空指针
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	pre := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}

	if list1 != nil {
		pre.Next = list1
	}

	if list2 != nil {
		pre.Next = list2
	}
	return head.Next
}
