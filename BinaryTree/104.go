package BinaryTree

var depth, res int

// 回溯:使用外部变量记录每个节点的最大深度,取最大值就可以得到最大深度
// depth在进入节点是添加深度,离开时减去深度
func maxDepth1(root *TreeNode) int {
	traverse(root)
	return res
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	depth++
	if root.Left == nil && root.Right == nil {
		// 到达叶子节点,更新最大深度
		if res < depth {
			res = depth
		}
	}
	traverse(root.Left)
	traverse(root.Right)
	depth--
}

// 动态规划:
// 把二叉树分成N个子树,求出左右子树的最大深度,返回其中最大的即可
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMax := maxDepth2(root.Left)
	rightMax := maxDepth2(root.Right)

	if leftMax < rightMax {
		return rightMax + 1
	}
	return leftMax + 1
}
