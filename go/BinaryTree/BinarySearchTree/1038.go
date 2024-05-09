package BinarySearchTree

func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	convertBST(root.Right)
	sum += root.Val
	root.Val = sum
	convertBST(root.Left)
	return root
}
