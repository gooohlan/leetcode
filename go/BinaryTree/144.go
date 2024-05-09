package BinaryTree

// 回溯
// 先记录当前值,再一次遍历左右节点,完成前序遍历
func preorderTraversal1(root *TreeNode) []int {
	var list []int
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		list = append(list, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return list
}

// 动态规划
// 拆分成左右节点的子问题,然后合并
func preorderTraversal2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	// 分解
	left := preorderTraversal2(root.Left)
	right := preorderTraversal2(root.Right)

	// 合并
	res = append(res, root.Val)
	res = append(res, left...)
	res = append(res, right...)
	return res
}
