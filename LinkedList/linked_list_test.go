package LinkedList

import (
	"fmt"
	"testing"
)

func PrintListNode(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func Test206(t *testing.T) {
	node := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	newNode := reverseList2(node)

	PrintListNode(newNode)
}

func TestReverseListN(t *testing.T) {
	node := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	node = reverseListN(node, 4)

	PrintListNode(node)
}
