package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := -1 << 30
	maxSum(root, &res)
	return res
}
func maxSum(root *TreeNode, maxValue *int) int {
	if root == nil {
		return 0
	}
	left := max(0, maxSum(root.Left, maxValue))
	right := max(0, maxSum(root.Right, maxValue))
	*maxValue = max(*maxValue, root.Val+left+right)
	return max(left, right) + root.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
