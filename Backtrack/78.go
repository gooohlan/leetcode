package Backtrack

// 最终生成的树大概是这样
//               [] 			s_0
//     [1]      [2]     [3]		s_1
// [1,2] [1,3] [2,3]			s_2
// [1,2,3]						s_3
// 首先生成元素个数为0的集合,就是空集,我们用s_0表示
// 依次向下推倒,注意看s_1 -> s_2, 为什么集合2只添加3?
// 因为集合不考虑顺序,如果我们添加了1,那么就会生成[2,1]的集合,与[1,2]重复,所以我们只考虑添加大于2的值
// 换句话说,我们通过保证元素之间的相对顺序不变来防止出现重复的子集
// 要求所有的集合,将整棵树的节点收集起来即可
func subsets(nums []int) [][]int {
	var res [][]int
	var track []int
	var backtrack func(start int)

	backtrack = func(start int) {
		item := make([]int, len(track))
		copy(item, track)
		res = append(res, item)

		// 回溯算法标准框架
		for i := start; i < len(nums); i++ {
			// 做选择
			track = append(track, nums[i])
			// 通过 start 参数控制树枝的遍历，避免产生重复的子集
			backtrack(i + 1)
			// 撤销选择
			track = track[:len(track)-1]
		}
	}

	backtrack(0)
	return res
}
