package BinaryTree

// 这道题是236的变体,不过也没有改变很多,把寻找两个变成了寻找N个
// 我们将节点存入map即可
func lowestCommonAncestor4(root *TreeNode, nodes []*TreeNode) *TreeNode {
	nodeMap := make(map[*TreeNode]bool) // 将需要查找的节点放入map
	for _, node := range nodes {
		nodeMap[node] = true
	}

	var find func(node *TreeNode, nodeMap map[*TreeNode]bool) *TreeNode

	find = func(node *TreeNode, nodeMap map[*TreeNode]bool) *TreeNode {
		if node == nil {
			return nil
		}

		if nodeMap[node] { // 找到既返回
			return node
		}
		left := find(node.Left, nodeMap)
		right := find(node.Right, nodeMap)
		if left != nil && right != nil {
			return node
		}
		if left != nil {
			return left
		}
		return right
	}

	return find(root, nodeMap)
}
