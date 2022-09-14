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

func TestPreorderTraversal(t *testing.T) {
	node := &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, &TreeNode{Val: 5}}},
		Right: &TreeNode{9, nil, nil},
	}
	fmt.Println(preorderTraversal1(node))
	fmt.Println(preorderTraversal2(node))
}
func TestDiameterOfBinaryTree(t *testing.T) {
	node := &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, &TreeNode{Val: 5}}},
		Right: &TreeNode{9, nil, nil},
	}
	fmt.Println(diameterOfBinaryTree(node))
}

func TestInvertTree(t *testing.T) {
	node := &TreeNode{
		Val:   4,
		Left:  &TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}},
		Right: &TreeNode{7, &TreeNode{Val: 6}, &TreeNode{Val: 9}},
	}
	node = invertTree1(node)
	node = invertTree2(node)

	fmt.Println(111)
}
