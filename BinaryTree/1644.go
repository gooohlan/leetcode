package BinaryTree

// 此题依旧是236的变体,此处不同的是,p和q可能会不存在于二叉树中
// 我们需要修改一下判断逻辑,如果二叉树中没有其中一个节点直接返回为空即可
// 而且因为PQ要同时存在,我们就不能遇到一个直接返回,需要把左右节点遍历后再判断
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var existP, existQ bool // 记录p,q是否存在

	var find func(root *TreeNode) *TreeNode
	find = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		left := find(node.Left)
		right := find(node.Right)
		if left != nil && right != nil { // 子节点找到目标值,直接返回
			return node
		}

		if node == p || node == q {
			if node == p {
				existP = true
			} else {
				existQ = true
			}
			return node
		}

		if left != nil {
			return left
		}
		return right
	}

	res := find(root)
	if !existQ || !existP {
		return nil
	}
	return res
}
