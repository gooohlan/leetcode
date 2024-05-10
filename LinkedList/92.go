package LinkedList

// 当m==1时,此题就变成了反转数组前N个元素
// 当m!=1时,我们将指针后移到开始节点,就又变成了反转前N个元素
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	// num := right - left + 1 // 需反转的数量
	prev := head
	for left > 2 {
		prev = prev.Next
		left--
		right--
	}
	prev.Next = reverseN(prev.Next, right-1) // right-1是因为prev.Next相当于进行了一次left--
	return head
}
