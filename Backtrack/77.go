package Backtrack

// 从78题的思路来看,相当于我们只求第二层的节点,所以添加 base case即可
func combine(n int, k int) [][]int {
	var (
		res       [][]int
		track     []int
		backtrack func(start int)
	)
	backtrack = func(start int) {
		if k == len(track) {
			// 遍历到了第 k 层，收集当前节点的值
			item := make([]int, len(track))
			copy(item, track)
			res = append(res, item)
			return
		}

		for i := start; i <= n; i++ {
			// 选择
			track = append(track, i)
			// 通过 start 参数控制树枝的遍历，避免产生重复的子集
			backtrack(i + 1)
			// 撤销选择
			track = track[:len(track)-1]
		}
	}
	backtrack(1)
	return res
}
