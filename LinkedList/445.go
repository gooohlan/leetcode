package LinkedList

func addTwoNumbers445(l1 *ListNode, l2 *ListNode) *ListNode {
    l1 = reverseList(l1)
    l2 = reverseList(l2)
    l3 := addTwoNumbers22(l1, l2)
    l3 = reverseList(l3)
    return l3
}
