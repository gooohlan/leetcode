package BinaryTree

// 遍历思路:
// 遍历每一个节点,在每个节点上交换自己的左右节点即可
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree1(root.Left)
	invertTree1(root.Right)
	return root
}

// 分解思路
// 每个子树完成对自己左子树的翻转,再完成右子树的翻转,然后交换左右子树的位置,即完成了此子树的翻转
// 整体的思路就是把对整棵树的翻转分解成了分别对左右子树的翻转,将左右子树返回结果进行处理,就变成了整棵树的翻转
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := invertTree2(root.Left)
	right := invertTree2(root.Right)

	root.Left = right
	root.Right = left
	return root
}
