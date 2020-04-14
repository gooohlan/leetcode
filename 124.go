package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := -1 << 30
	maxSum(root, &res)
	return res
}
func maxSum(root *TreeNode, maxValue *int) int{
	if root == nil {
		return 0
	}
	left := max(0, maxSum(root.Left, maxValue))
	right := max (0, maxSum(root.Right, maxValue))
	*maxValue = max(*maxValue, root.Val + left + right)
	return max(left, right) + root.Val
}


func max(a, b int) int {
	if a > b{
		return a
	}
	return b
}

func main() {
	root := &TreeNode{Val:-10}
	root.Left = &TreeNode{Left:nil, Val:9, Right:nil}
	root.Right = &TreeNode{Left: &TreeNode{Val:15}, Val:20, Right:&TreeNode{Val:7}}

	fmt.Println(maxPathSum(root))
}