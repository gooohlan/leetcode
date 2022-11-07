package BinarySearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉搜索树的中序遍历是升序排列的,此题需要找到第N个,我们直接中序排列取出即可
func kthSmallest(root *TreeNode, k int) int {
	var res int

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		k--
		if k == 0 {
			res = node.Val
			return
		}
		dfs(node.Right)
	}

	dfs(root)
	return res
}

type MyBst struct {
	root    *TreeNode
	nodeNum map[*TreeNode]int
}

// 统计node子节点个数
func (m *MyBst) countNodeNum(node *TreeNode) int {
	if node == nil {
		return 0
	}

	m.nodeNum[node] = m.countNodeNum(node.Left) + m.countNodeNum(node.Right) + 1
	return m.nodeNum[node]
}

func (m *MyBst) kthSmallest(k int) int {
	node := m.root

	for {
		leftNodeNum := m.nodeNum[node.Left]
		if leftNodeNum < k-1 { // 左侧不够,去右边找
			node = node.Right
			k -= leftNodeNum + 1
		} else if leftNodeNum == k-1 {
			return node.Val
		} else {
			node = node.Left
		}
	}
}

func kthSmallest2(root *TreeNode, k int) int {
	t := &MyBst{root, map[*TreeNode]int{}}
	t.countNodeNum(root)
	return t.kthSmallest(k)
}
