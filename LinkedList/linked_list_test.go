package linked_list

import (
	"fmt"
	"testing"
)

func NewListNode(nums ...int) *ListNode {
	head := &ListNode{}
	pre := head
	for _, num := range nums {
		node := &ListNode{Val: num}
		pre.Next = node
		pre = pre.Next
	}
	return head.Next
}

func TestMergeTwoLists(t *testing.T) {
	l1 := NewListNode(1, 4, 7)
	l2 := NewListNode(2, 3, 5, 7, 8)

	l3 := mergeTwoLists(l1, l2)
	fmt.Println(l3)
}
