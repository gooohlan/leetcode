package Backtrack

func combinationSum(candidates []int, target int) [][]int {
	var (
		res       [][]int
		track     []int
		trackSum  int // 记录track 之和
		backtrack func(start int)
	)

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
			track = append(track, candidates[i])
			trackSum += candidates[i]
			backtrack(i)
			trackSum -= candidates[i]
			track = track[:len(track)-1]
		}
	}

	backtrack(0)
	return res
}
