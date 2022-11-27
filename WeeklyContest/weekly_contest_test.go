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

	// fmt.Println(removeNodes(head))
	head = removeNodes2(head)
	fmt.Println(1)
}

func Test6248(t *testing.T) {
	nums := []int{5, 1, 7, 4, 2, 6, 3}
	fmt.Println(countSubarrays(nums, 4))
}
