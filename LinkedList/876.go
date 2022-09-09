package linked_list

// 要找到中间节点,常规的办法是遍历一遍,拿到链表的长度,再遍历一次得到n/2个节点,也就是中间节点
// 使用快慢指针的方式,申明两个指针,每当慢指针前进一步,快指针就前进两步,这当看指针走到链表末尾,慢指针刚好到链表的中间节点
// 当链表长度为偶数时,快指针走到指针末尾,而慢指针刚好走到两个中间节点靠后的那个节点,符合题意
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
