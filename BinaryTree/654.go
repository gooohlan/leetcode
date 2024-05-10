package BinaryTree

// 遍历数组找出最大的数作为父节点,然后递归最大数左边和右边的数进行递归构建,作为父节点的左右子树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxIndex := 0
	for i, num := range nums {
		if num > nums[maxIndex] {
			maxIndex = i
		}
	}

	return &TreeNode{
		Val:   nums[maxIndex],
		Left:  constructMaximumBinaryTree(nums[:maxIndex]),
		Right: constructMaximumBinaryTree(nums[maxIndex+1:]),
	}
}

// 单调栈
// 遍历数组,与栈顶进行比较,
// 大于栈顶则让栈顶出栈,加入自己的左子树
// 小于栈顶则入栈,自己加入栈顶右子树
func constructMaximumBinaryTree2(nums []int) *TreeNode {
	tree := make([]*TreeNode, len(nums))
	var stack []int // 数组下标存入栈,方便比较
	for i, num := range nums {
		tree[i] = &TreeNode{Val: num}
		// 大于栈顶
		for len(stack) > 0 && num > nums[stack[len(stack)-1]] {
			tree[i].Left = tree[stack[len(stack)-1]] // 栈顶出栈加入左子树
			stack = stack[:len(stack)-1]
		}

		// 小于栈顶,入栈,自己加入栈顶右子树
		if len(stack) > 0 {
			tree[stack[len(stack)-1]].Right = tree[i]
		}
		stack = append(stack, i)
	}
	return tree[stack[0]]
}
