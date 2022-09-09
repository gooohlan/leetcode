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

func TestPartition(t *testing.T) {
	head := NewListNode(1, 4, 3, 2, 5, 2)
	node := partition(head, 3)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func TestMergeKLists(t *testing.T) {
	l1 := NewListNode(1, 4, 5)
	l2 := NewListNode(1, 3, 4)
	l3 := NewListNode(2, 6)

	newListNode := mergeKLists1([]*ListNode{l1, l2, l3})
	for newListNode != nil {
		fmt.Println(newListNode.Val)
		newListNode = newListNode.Next
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	l := NewListNode(1, 2, 3, 4, 5)
	newListNode := removeNthFromEnd(l, 2)
	for newListNode != nil {
		fmt.Println(newListNode.Val)
		newListNode = newListNode.Next
	}
}

func TestMiddleNode(t *testing.T) {
	l := NewListNode(1, 2, 3, 4, 5)
	mid := middleNode(l)
	for mid != nil {
		fmt.Println(mid.Val)
		mid = mid.Next
	}
}
