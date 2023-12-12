package WeeklyContest

func reverseList(head *ListNode) *ListNode {
    var pre *ListNode
    cur := head
    for cur != nil {
        // curr.Next, prev, curr = prev, curr, curr.Next
        cur, cur.Next, pre = cur.Next, pre, cur
    }
    return pre
}
func doubleIt(head *ListNode) *ListNode {

}
