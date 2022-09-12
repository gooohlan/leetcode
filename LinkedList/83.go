package linked_list

// 此题与26题删除有序数组中的重复项思路一致,唯一的区别就是把数组的赋值操作变成操作指针而已
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil {
		if slow.Val != fast.Val {
			// nums[slow] = nums[fast]
			slow.Next = fast
			// slow ++
			slow = slow.Next
		}
		// fast ++
		fast = fast.Next
	}
	// 断开与后面重复元素的链接
	slow.Next = nil
	return head
}
