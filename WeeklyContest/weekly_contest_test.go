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

func Test6255(t *testing.T) {
	roads := [][]int{
		[]int{1, 2, 9},
		[]int{2, 3, 6},
		[]int{2, 4, 5},
		[]int{1, 4, 7},
	}

	fmt.Println(minScore(4, roads))
}

func Test6256(t *testing.T) {
	edges := [][]int{
		[]int{1, 2},
		[]int{1, 4},
		[]int{1, 5},
		[]int{2, 6},
		[]int{2, 3},
		[]int{4, 6},
	}
	fmt.Println(magnificentSets(6, edges))
}
