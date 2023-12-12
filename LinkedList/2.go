package LinkedList

func addTwoNumbers2_1(l1 *ListNode, l2 *ListNode) *ListNode {
    return addTwo(l1, l2, 0)
}

func addTwo(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
    if l1 == nil && l2 == nil {
        if carry != 0 {
            return &ListNode{Val: carry}
        }
        return nil
    }
    if l1 == nil { // 保证l1不为空
        l1, l2 = l2, l1
    }
    carry += l1.Val
    if l2 != nil {
        carry += l2.Val
        l2 = l2.Next
    }
    l1.Val = carry % 10
    l1.Next = addTwo(l1.Next, l2, carry/10)
    return l1
}

func addTwoNumbers22(l1 *ListNode, l2 *ListNode) *ListNode {
    pre := &ListNode{}
    cur := pre
    carry := 0
    for l1 != nil || l2 != nil || carry != 0 { // 有一个节点不为空或者还有进位,继续循环
        if l1 != nil {
            carry += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            carry += l2.Val
            l2 = l2.Next
        }
        
        cur.Next = &ListNode{Val: carry % 10}
        carry /= 10
        cur = cur.Next
    }
    return pre.Next
}
