package linked_list

// 定义快慢指针,慢指针每次走1步,快指针走两步
// 如果链表中存在环,那么快指针会追上慢指针,否则快指针会走到链表末尾
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
