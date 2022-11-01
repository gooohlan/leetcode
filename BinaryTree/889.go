package BinaryTree

// 前序遍历: [ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
// 后续遍历: [[左子树的前序遍历结果], [右子树的前序遍历结果] 根节点 ]
// 前面都是在中序遍历中去寻找中心节点,这里稍微有所改动
// 我们将前序遍历视为标准,那么前序遍历中的第二个数组,就是下一次的中心节点
// 那么在后续遍历数组中,下次的中心节点前的即为前序遍历部分,后面的则为后续遍历部分
// 以后续遍历为标准以此类推
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	if len(preorder) == 1 {
		return root
	}
	i := 0
	for ; i < len(postorder); i++ {
		if preorder[1] == postorder[i] { // 以左序遍历为根节点时,在后续节点中找到根节点
			break
		}
	}
	root.Left = constructFromPrePost(preorder[1:len(postorder[:i+1])+1], postorder[:i+1])
	root.Right = constructFromPrePost(preorder[len(postorder[:i+1])+1:], postorder[i+1:len(postorder)-1])
	return root
}
