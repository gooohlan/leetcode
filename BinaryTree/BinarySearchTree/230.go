package BinarySearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉搜索树的中序遍历是升序排列的,此题需要找到第N个,我们直接中序排列取出即可
func kthSmallest(root *TreeNode, k int) int {
	var res int

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		k--
		if k == 0 {
			res = node.Val
			return
		}
		traverse(node.Right)
	}

	traverse(root)
	return res
}
