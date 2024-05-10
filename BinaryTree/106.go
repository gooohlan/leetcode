package BinaryTree

// 参照上一题的思路
// 后续遍历: [[左子树的前序遍历结果], [右子树的前序遍历结果] 根节点 ]
// 中序遍历 [[左子树的前序遍历结果], 根节点, [右子树的前序遍历结果] ]
func buildTree106(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{postorder[len(postorder)-1], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] { // 找到中序遍历根节点
			break
		}
	}

	root.Left = buildTree106(inorder[:i], postorder[:i])
	root.Right = buildTree106(inorder[i+1:], postorder[i:len(postorder)-1])
	return root
}
