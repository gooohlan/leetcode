package BinarySearchTree

// 因为二叉搜索树本身就是有序的,它的中序遍历是升序的
// 根据二叉搜索树的特性,左子树小于父节点小于右子树,我们先遍历右子树,最后遍历左子树不就成了降序吗,这时累加即可
var sum int

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	convertBST(root.Right)
	sum += root.Val
	root.Val = sum
	convertBST(root.Left)
	return root
}
