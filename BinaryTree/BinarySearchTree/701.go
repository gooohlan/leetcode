package BinarySearchTree

// 跟二叉搜索树一样,现在找到指定节点位置,再插入,这里不同的是需要改变二叉树的结构
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
