package BFS

import (
	"fmt"
	"testing"
)

func TestMinDepth(t *testing.T) {
	node := &TreeNode{
		1,
		&TreeNode{2, &TreeNode{4, nil, nil}, nil},
		&TreeNode{3, nil, &TreeNode{5, nil, nil}},
	}
	depth := minDepth(node)
	fmt.Println(depth)
}

func TestOpenLock(t *testing.T) {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	fmt.Println(openLock(deadends, "0202"))
}
