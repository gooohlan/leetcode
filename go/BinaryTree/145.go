package BinaryTree

// 回溯
// 代码与144几乎一模一样,不同的只是在添加节点时位置发生了变化
func postorderTraversal(root *TreeNode) []int {
	var list []int
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		preorder(node.Left)
		preorder(node.Right)
		list = append(list, node.Val)
	}
	preorder(root)
	return list
}

func postorderTraversal2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	// 分解
	left := postorderTraversal2(root.Left)
	right := postorderTraversal2(root.Right)

	// 合并
	res = append(res, left...)
	res = append(res, right...)
	res = append(res, root.Val)
	return res
}
