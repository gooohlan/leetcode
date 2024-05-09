package BinaryTree

var depth, res int

// 回溯:使用外部变量记录每个节点的最大深度,取最大值就可以得到最大深度
// depth在进入节点是添加深度,离开时减去深度
// leetcode不支持全局变量,所以要把traverse写成定义为局部变量的方式进行调用
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

// 函数声明为局部变量调用方式
func maxDepth(root *TreeNode) int {
	var traverse func(root *TreeNode)
	var depth int
	var result int
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 前序位置
		depth++
		// 子节点
		if root.Left == nil && root.Right == nil {
			if depth > result {
				result = depth
			}
		}
		traverse(root.Left)
		traverse(root.Right)
		// 后序位置
		depth--
	}
	traverse(root)
	return result
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
