package BinaryTree

// 嵌套的整型列表可以看做一个N叉树,树上的叶子节点对应一个整数，非叶节点对应一个列表
// 在这棵树上深度优先搜索的顺序就是迭代器遍历的顺序。
// 我们可以先遍历整个嵌套列表，将所有整数存入一个数组，然后遍历该数组从而实现next和hasNext方法。
type NestedIterator struct {
	nums []int
}

type NestedInteger struct {
}

// 自带API部分不实现
func (n NestedInteger) IsInteger() bool {
	return true
}

func (n NestedInteger) GetInteger() int {
	return 0
}

func (n NestedInteger) GetList() []*NestedInteger {
	return nil
}

func Constructor314(nestedList []*NestedInteger) *NestedIterator {
	var nums []int
	var dfs func([]*NestedInteger)

	dfs = func(nestedList []*NestedInteger) {
		for _, nest := range nestedList {
			if nest.IsInteger() { // 如果是一个整数,添加到到nums里
				nums = append(nums, nest.GetInteger())
			} else {
				dfs(nest.GetList())
			}
		}
	}

	dfs(nestedList)
	return &NestedIterator{
		nums: nums,
	}
}

func (n *NestedIterator) Next() int {
	num := n.nums[0]
	n.nums = n.nums[1:]
	return num
}

func (n *NestedIterator) HasNext() bool {
	return len(n.nums) > 0
}

// 栈操作
type NestedIterator2 struct {
	Stack []*NestedInteger // 栈是用来暂存元素的
}

func Constructor314_2(nestedList []*NestedInteger) *NestedIterator2 {
	stack := []*NestedInteger{}
	for i := 0; i < len(nestedList); i++ {
		stack = append(stack, nestedList[i]) // 顺序入栈, 出栈时送数组最前取即可
	}
	return &NestedIterator2{Stack: stack}
}

func (n *NestedIterator2) Next() int {
	top := n.Stack[0]     // 获取栈顶
	n.Stack = n.Stack[1:] // 移除栈顶
	return top.GetInteger()
}

// 因为调用Next前总是调用HasNext, 调用HasNext时把栈顶替换为integer
func (n *NestedIterator2) HasNext() bool {
	for len(n.Stack) > 0 { // 直到栈为空
		top := n.Stack[0]    // 获取栈顶
		if top.IsInteger() { // 如果栈顶为inter,什么也不做并返回true
			return true
		}

		n.Stack = n.Stack[1:] // 栈顶为list,先弹出栈顶
		list := top.GetList() // 获取list,从新组成栈顶放入栈
		var newTop []*NestedInteger
		for _, integer := range list {
			newTop = append(newTop, integer)
		}
		n.Stack = append(newTop, n.Stack...) // 从顶部入栈,这里利用了切片的特性,而不用像真正的栈那样后进先出
	}
	return false // 栈为空
}
