package LinkedList

import (
	"fmt"
	"testing"
)

func PrintListNode(node *ListNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	PrintListNode(node.Next)
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

func Test92(t *testing.T) {
	node := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}}
	node = reverseBetween(node, 4, 6)

	PrintListNode(node)
}

func Test25(t *testing.T) {
	node := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}}
	node = reverseKGroup(node, 4)
	PrintListNode(node)
}

func Test234(t *testing.T) {
	node := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}
	fmt.Println(isPalindrome2(node))
}
