package BinaryTree

// 因为需要再原二叉树上进行修改,所以没法使用遍历的方式
// 使用分解问题的方式
// 对于每个节点,先将他的左右节点拉平
// 再将右子树接到左子树下方, 把左子树变成右子树,自此这个节点的二叉树转链表完成
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// 先拉平左右子树
	flatten(root.Left)
	flatten(root.Right)

	// 取出已经拉平的左右子树
	left := root.Left
	right := root.Right

	// 先将左子树变成右子树
	root.Left = nil
	root.Right = left

	// 将原有右子树连接到新右子树下方
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}
