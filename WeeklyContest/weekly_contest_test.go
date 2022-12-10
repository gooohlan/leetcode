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

func Test6261(t *testing.T) {
	strs := []string{"1", "01", "001", "0001"}
	fmt.Println(maximumValue(strs))
}

func Test6262(t *testing.T) {
	vals := []int{
		17, 49, -34, -17, -7, -23, 24,
	}
	edges := [][]int{
		[]int{3, 1},
		[]int{5, 1},
		[]int{0, 3},
		[]int{4, 6},
		[]int{1, 4},
		[]int{3, 4},
		[]int{6, 3},
		[]int{2, 6},
		[]int{5, 2},
		[]int{1, 6},
		[]int{6, 0},
		[]int{2, 3},
		[]int{3, 5},
		[]int{2, 1},
		[]int{0, 2},
		[]int{5, 0},
		[]int{0, 4},
	}
	fmt.Println(maxStarSum(vals, edges, 6))
}
