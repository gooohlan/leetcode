package BinaryTree

import "math"

// 因为完全二叉树中存在满二叉树,对于满二叉树,节点个数与层数成指数关系
// 一个完全二叉树中肯定存在满二叉树和完全二叉树的组合,满二叉树我们使用指数求结果,完全二叉树就按照正常逻辑求
func countNodes(root *TreeNode) int {
	l, r := root, root
	hl, hr := 0, 0
	for l != nil {
		l = l.Left
		hl++
	}

	for r != nil {
		r = r.Right
		hr++
	}
	if hl == hr { // 完全二叉树
		return int(math.Pow(2, float64(hl))) - 1
	}
	// 非完全二叉树,按照传统求法计算
	left := countNodes(root.Left)
	right := countNodes(root.Right)
	return 1 + left + right
}
