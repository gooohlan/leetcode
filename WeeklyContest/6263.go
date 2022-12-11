package WeeklyContest

// 来回不能跳同一块石头(首尾不算),假设一共8块
// 来: 1, 3, 5, 7
// 回: 2, 4, 6
// 则求出这两两相邻的节点中最大值即可
func maxJump(stones []int) int {
	res := stones[1]
	for i := 2; i < len(stones); i++ {
		res = max(res, stones[i]-stones[i-2])
	}

	return res
}

// 合并一下来回
