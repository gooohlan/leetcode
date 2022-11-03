package BinaryTree

// 与层级遍历不同,我们不需要出队操作,而是额外声明一个变量存储下级节点,没循环完一级,替换掉主队列
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var queue []*TreeNode

	queue = append(queue, root)
	for i := 0; len(queue) != 0; i++ {
		res = append(res, []int{})

		var p []*TreeNode
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			res[i] = append(res[i], node.Val)

			// 下一级节点加入P
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		// 替换Q的值
		queue = p
	}
	return res
}
