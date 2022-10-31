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
	// fmt.Println(preorderTraversal2(node))
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

func TestConnect(t *testing.T) {
	node := &Node{
		1,
		&Node{
			2,
			&Node{4, &Node{8, nil, nil, nil}, &Node{9, nil, nil, nil}, nil},
			&Node{5, &Node{10, nil, nil, nil}, &Node{11, nil, nil, nil}, nil},
			nil,
		},
		&Node{
			3,
			&Node{6, &Node{12, nil, nil, nil}, &Node{13, nil, nil, nil}, nil},
			&Node{7, &Node{14, nil, nil, nil}, &Node{15, nil, nil, nil}, nil},
			nil,
		},
		nil,
	}

	connect(node)
}

func Test114(t *testing.T) {
	node := &TreeNode{
		1,
		&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{5, &TreeNode{6, nil, nil}, nil},
	}

	flatten(node)
	fmt.Println(node)
}

func Test654(t *testing.T) {
	nums := []int{3, 2, 1, 6, 0, 5}
	node := constructMaximumBinaryTree2(nums)
	fmt.Println(node)
}

func Test105(t *testing.T) {
	preorder := []int{1, 2}
	inorder := []int{2, 1}
	node := buildTree(preorder, inorder)
	fmt.Println(node)
}
