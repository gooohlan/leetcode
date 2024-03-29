package linked_list

// 需要判断两个链表是否有交点,其核心在于两个链表同时到达相交点c1
// 先让两个指针同时走到距离尾部相同的距离
// 然后依次往后遍历, 如果p1与p2相等,他们就走到了相交的节点,否则会走到链表尾部
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	var lenA, lenB int
	for p := headA; p != nil; p = p.Next {
		lenA++
	}
	for p := headB; p != nil; p = p.Next {
		lenB++
	}
	p1, p2 := headA, headB
	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			p1 = p1.Next
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			p2 = p2.Next
		}
	}

	// 注意题意,提供的两个链表中存在相同链表,所以直接判断两个链表相等即可
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}

// 核心依旧在于同事到达相交点c1
// 我们让两个链表收尾相连,这样两个链表的长度一致,即p1 = headA + headB p2 = headB + headA
// 这样我们同时遍历两个指针,就可以同时到底c1,如果没有相交点,也会同时走到指针的末尾
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		}
		if p2 == nil {
			p2 = headA
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
