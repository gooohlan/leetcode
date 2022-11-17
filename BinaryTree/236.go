package BinaryTree

// 两个节点同时寻找,此时分两种情况
// 当前节点是其中一个,直接返回即可
// 子节点找到两个值之一,那么此节点就是公共父节点
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 找到直接返回
	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
