package DP

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 使用一个长度为2的数组表示选择与不选择
// arr[0]表示此节点不抢最大值
// arr[1]表示此节点抢的结果
func rob337(root *TreeNode) int {
	arr := dfs(root)
	return max(arr[0], arr[1])
}

func dfs(root *TreeNode) (arr [2]int) {
	if root == nil {
		return arr
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	// arr[0] 表示当前节点不抢,子节点抢, 从左右节结果中取出最大的相加
	arr[0] = max(left[0], left[1]) + max(right[0], right[1])
	// arr[1] 表示当前节点抢,子节点不抢
	arr[1] = root.Val + left[0] + right[0]
	return arr
}
