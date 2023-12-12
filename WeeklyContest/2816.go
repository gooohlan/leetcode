package WeeklyContest

// 方法 1: 直接两个链表相加,见 445
// 方法 2: 每个数 *2 再考虑进位

func doubleIt(head *ListNode) *ListNode {
    if head.Val > 4 {
        head = &ListNode{0, head}
    }
    
    for cur := head; cur != nil; cur = cur.Next {
        cur.Val = cur.Val * 2 % 10
        if cur.Next != nil && cur.Next.Val > 4 {
            cur.Val++
        }
    }
    return head
}
