package Backtrack

import "sort"

// 这也是一个组合问题,只不过是由之前的求第N层变成了和为target
// 我们使用一个变量记录回溯元素上的和, 改一下base case即可
func combinationSum2(candidates []int, target int) [][]int {
	var (
		res       [][]int
		track     []int
		trackSum  int // 记录track 之和
		backtrack func(start int)
	)

	sort.Ints(candidates)
	backtrack = func(start int) {
		if trackSum == target {
			item := make([]int, len(track))
			copy(item, track)
			res = append(res, item)
			return
		}

		// 超过目标和直接跳过
		if trackSum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			// 剪枝处理，值相同的相邻树枝，只遍历第一条
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			trackSum += candidates[i]
			track = append(track, candidates[i])
			backtrack(i + 1)
			track = track[:len(track)-1]
			trackSum -= candidates[i]
		}
	}

	backtrack(0)
	return res
}
