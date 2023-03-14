package LinkedList

import "math/rand"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
    head *ListNode
}

func Constructor(head *ListNode) Solution {
    return Solution{head}
}

func (this *Solution) GetRandom() int {
    res, i := 0, 1
    node := this.head
    for node != nil {
        if rand.Intn(i) == 0 {
            res = node.Val
        }
        i++
        node = node.Next
    }
    return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
