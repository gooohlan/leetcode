package BinaryTree

// 此题比236更加简单,因为树是一个二叉搜索树,利用二叉搜索树的特性可得
// p < root.val < q时,root就是pq的公共祖先
// 如果p > root.val 那么就去右子树寻找, q > root.val就去左子树寻找
func lowestCommonAncestor235(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val { // 保证p < q
		p, q = q, p
	}

	var find func(*TreeNode) *TreeNode
	find = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		if node.Val < p.Val {
			return find(node.Right)
		}

		if node.Val > q.Val {
			return find(node.Left)
		}

		return node
	}

	return find(root)
}
