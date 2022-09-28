package BFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 从上到下,从左到右进行搜索,第一个发现的叶子节点,就是最小深度所在
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	depth := 1
	queue = append(queue, root)

	// 从上到下
	for len(queue) != 0 {
		// 从左到右
		sz := len(queue)
		for i := 0; i < sz; i++ {
			node := queue[0]
			queue = queue[1:] // 模拟出队操作
			// 判断是否到达终点
			if node.Left == nil && node.Right == nil {
				return depth
			}
			// 将左右节点加入队列中
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}
