package linked_list

func partition(head *ListNode, x int) *ListNode {
	head1 := &ListNode{} // 存放小于x的链表
	head2 := &ListNode{} // 存放大于等于x的链表
	p1, p2 := head1, head2
	for head != nil {
		if head.Val < x {
			p1.Next = head
			p1 = p1.Next
		} else {
			p2.Next = head
			p2 = p2.Next
		}
		head = head.Next
	}
	p2.Next = nil
	// 连接两个链表
	p1.Next = head2.Next
	return head1.Next
}
