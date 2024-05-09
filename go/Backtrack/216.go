package Backtrack

// 元素无重不可复选的组合,多了和为n个数为k的限制
func combinationSum3(k int, n int) [][]int {
	var (
		res       [][]int
		track     []int
		trackSum  int
		backtrack func(start int)
	)

	backtrack = func(start int) {
		if len(track) == k && trackSum == n {
			item := make([]int, len(track))
			copy(item, track)
			res = append(res, item)
			return
		}

		if trackSum > n || len(track) > k {
			return
		}

		for i := start; i < 10; i++ {
			track = append(track, i)
			trackSum += i
			backtrack(i + 1)
			trackSum -= i
			track = track[:len(track)-1]
		}
	}

	backtrack(1)
	return res
}
