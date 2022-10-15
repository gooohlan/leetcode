package linked_list

// 假设快慢指针相遇,慢指针走了k步,那么快指针一定走了2k步,快指针多走的这k步就是在环里转圈圈,所以k的值就是环长度的整数倍
// 假设相遇点为m,那么环起点就为k-m,也就是说,从链表头部前进k-m就可以到达环起点
// 因为相遇点为m,k为环长度的倍数,所以只需要再走k-m步就可以到底环起点
// 所以当我们相遇时,重新定义一个指针从头出发,与慢指针再次相遇的地方就是环起点
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
