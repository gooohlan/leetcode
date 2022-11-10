package BinarySearchTree

// 删除时分三种情况
// 1 此节点无子节点,直接删除即可
// 2 此节点有一个非空子节点,那么用子节点替换即可
// 3 有两个子节点，麻烦了，为了不破坏 BST 的性质，必须找到左子树中最大的那个节点，或者右子树中最小的那个节点来接替自己
func deleteNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		// 处理1和2
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		// 情况3 获取右边最小的
		minNode := getMin(root.Right)
		// 删除右子树最小节点
		root.Right = deleteNode(root.Right, minNode.Val)
		// 替换当前节点
		minNode.Left = root.Left
		minNode.Right = root.Right
		root = minNode
	}

	if root.Val > val {
		root.Left = deleteNode(root.Left, val)
	}
	if root.Val < val {
		root.Right = deleteNode(root.Right, val)
	}
	return root
}

// 找到最小的节点
func getMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
