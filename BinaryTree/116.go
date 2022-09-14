package BinaryTree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 因为是完美二叉树,所以左右节点都存在,连接左右子树简单,难点在于连接左子树的右子树与右子树的左子树
// 定义一个函数,同时数的左右子树,分别连接他们各自子树的左右子树,在完成左子树的右子树与右子树的左子树连接
func connect(root *Node) *Node {
	connectTraverse(root.Left, root.Right)
	return root
}

func connectTraverse(leftNode *Node, rightNode *Node) {
	if leftNode == nil || rightNode == nil {
		return
	}
	// 首先连接自身
	leftNode.Next = rightNode

	// 连接想通父节点的两个子节点
	connectTraverse(leftNode.Left, leftNode.Right)
	connectTraverse(rightNode.Left, rightNode.Right)
	// 连接相邻节点的子节点
	connectTraverse(leftNode.Right, rightNode.Left)
}
