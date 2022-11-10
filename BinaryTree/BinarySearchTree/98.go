package BinarySearchTree

// 限定以 root 为根的子树节点必须满足 max.val > root.val > min.val
func isValidBST(root *TreeNode) bool {
	var bst func(node, min, max *TreeNode) bool

	bst = func(node, min, max *TreeNode) bool {
		if node == nil {
			return true
		}

		if min != nil && min.Val >= node.Val {
			return false
		}
		if max != nil && max.Val < node.Val {
			return false
		}

		// 限定左子树的最大值是 node.val，右子树的最小值是 node.val
		return bst(node.Left, min, node) && bst(node.Right, node, max)
	}
	return bst(root, nil, nil)
}
