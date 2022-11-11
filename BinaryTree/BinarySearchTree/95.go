package BinarySearchTree

// 参考95题的暴力遍历
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	return build(1, n)
}

func build(lo, hi int) []*TreeNode {
	if lo > hi {
		return []*TreeNode{nil}
	}

	var res []*TreeNode

	for i := lo; i <= hi; i++ {
		// 找到所有合法的左右子树
		leftTree := build(lo, i-1)
		rightTree := build(i+1, hi)
		// 将左右子树组合成二叉搜索树
		for _, left := range leftTree {
			for _, right := range rightTree {
				tree := &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				}
				res = append(res, tree)
			}
		}
	}
	return res
}
