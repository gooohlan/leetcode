package BinaryTree

// 前序遍历: [ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
// 中序遍历 [[左子树的前序遍历结果], 根节点, [右子树的前序遍历结果] ]
// 根据前序遍历即可找出根节点,在拆分左右子树遍历结果去递归,即可得到整棵树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] { // 找到中序遍历根节点
			break
		}
	}

	// 无论怎么遍历,同边节点个数始终一样
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i]) // len(inorder[:i])+1防止越界
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
