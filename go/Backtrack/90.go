package Backtrack

import "sort"

// 与前面的组合问题思路大概一致,不同点在于有重复元素,但是不能包含重复子集
// 我们将原数组排序,那么在每一层循环遍历时,遇到已经遍历过得则跳过,这样就解决了重复子元素的问题
func subsetsWithDup(nums []int) [][]int {
	var (
		res       [][]int
		track     []int
		backtrack func(start int)
	)
	sort.Ints(nums)

	backtrack = func(start int) {
		item := make([]int, len(track))
		copy(item, track)
		res = append(res, item)
		for i := start; i < len(nums); i++ {
			// 剪枝处理，值相同的相邻树枝，只遍历第一条
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			track = append(track, nums[i])
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}

	backtrack(0)
	return res
}
