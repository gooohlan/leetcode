package linked_list

// 假设链表有n个节点,那么倒数第K个节点,就是正数第n-k+1个节点
// 遍历一次链表,获取链表长度n,再重新遍历一次找到第n-k+1个节点即可
// 能不能遍历一次就搞定呢
// 首先我们定义一个指针p1,p1先走k步,那么他还需要走n-k步即可到达尾部
// 此时我们再申明一个指针p2,指向头部, p1和p2同时前进,当p1到达尾部时,p1正好走了n-k步,在n-k+1个节点上,就是倒数第k个节点
// 为了方便删除操作,申明p2时申明一个前置节点,按照上述步骤,p1到达末尾, p2到达倒数第N个节点的前置节点, 直接删除p2下一个节点即可

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := &ListNode{0, head}
	p1, p2 := head, p
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p2.Next = p2.Next.Next
	return p.Next
}
