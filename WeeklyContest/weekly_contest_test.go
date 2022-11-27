package WeeklyContest

import (
	"fmt"
	"testing"
)

func Test6245(t *testing.T) {
	fmt.Println(pivotInteger(1))
}

func Test6246(t *testing.T) {
	fmt.Println(appendCharacters("abcde", "a"))
}

func Test6247(t *testing.T) {
	head := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 13,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
	}

	fmt.Println(removeNodes(head))
}
