package BinaryTree

// 使用分解法
// 定义diameter
// 分别求出左右子树的最大深度,和为该节点最大直径
// 返回该节点最大深度,以便父节点使用
var maxDiameter int

func diameterOfBinaryTree(root *TreeNode) int {
	return diameter(root)
}

func diameter(node *TreeNode) int {
	if node == nil {
		return 0
	}
	//
	leftDiameter := diameter(node.Left)
	rightDiameter := diameter(node.Right)

	diameter := leftDiameter + rightDiameter
	maxDiameter = max(maxDiameter, diameter)

	if leftDiameter > rightDiameter {
		return leftDiameter + 1
	}
	return max(leftDiameter, rightDiameter) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
