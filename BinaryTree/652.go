package BinaryTree

import "fmt"

// 通过前面用到的序列化技巧,序列化当前树,存入map中记录出现次数
// 最后从map中取出出现次数大于1的
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	repeat := make(map[*TreeNode]struct{}) // 记录重复出现的数
	seen := make(map[string]*TreeNode)     // 定义序列化与树的映射关系
	var dfs func(*TreeNode) string

	dfs = func(node *TreeNode) string {
		if node == nil {
			return "-"
		}

		left := dfs(node.Left)
		right := dfs(node.Right)
		subTree := fmt.Sprintf("%s,%s,%d", left, right, node.Val)
		if n, ok := seen[subTree]; ok {
			repeat[n] = struct{}{}
		} else {
			seen[subTree] = node
		}
		return subTree
	}

	dfs(root)
	res := make([]*TreeNode, 0, len(repeat))
	for treeNode := range repeat {
		res = append(res, treeNode)
	}
	return res
}
