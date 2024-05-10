package LinkedList

var left *ListNode

// 利用递归的方法倒序取出链表中的节点, 与正序中的一一比较即可
// 递归的方式相当于把节点放入了一个栈
func isPalindrome(head *ListNode) bool {
	left = head
	return traverse(head)
}

func traverse(right *ListNode) bool {
	if right == nil {
		return true
	}
	if !traverse(right.Next) {
		return false
	}
	if right.Val != left.Val {
		return false
	}
	left = left.Next
	return true
}

// 利用双指针先找到中心节点,然后反转中心节点后面部分, 与前面进行对比
// left可以省去
func isPalindrome2(head *ListNode) bool {
	slow, fast, left := head, head, head
	for fast.Next != nil && fast.Next.Next != nil { // 长度为奇数时slow会指向中心节点, 偶数会指向中心节点前一位
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 反转后半部分
	right := reverseList(slow.Next)
	defer func() {
		reverseList(slow.Next) // 还原链表
	}()
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}
