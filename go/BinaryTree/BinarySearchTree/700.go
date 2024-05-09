package BinarySearchTree

// 利用二叉搜索树的特性,比较左右节点后去指定节点寻找
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	return root
}
