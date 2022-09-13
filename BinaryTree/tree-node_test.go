package BinaryTree

import (
	"fmt"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	node := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: &TreeNode{2, nil, nil},
	}
	fmt.Println(maxDepth1(node))
	fmt.Println(maxDepth2(node))
}
