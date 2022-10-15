package LinkedList

import (
	"fmt"
	"testing"
)

func Test206(t *testing.T) {
	node := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	newNode := reverseList2(node)

	for newNode != nil {
		fmt.Println(newNode.Val)
		newNode = newNode.Next
	}
}
