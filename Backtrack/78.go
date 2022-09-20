package Backtrack

func subsets(nums []int) [][]int {
	var res [][]int
	var track []int
	backtrack78(&res, track, nums, 0)
	return res
}

// 多申明了两个变量,避免使用全局变量,leetcode会报错
func backtrack78(res *[][]int, track, nums []int, index int) {
	// 前序位置，每个节点的值都是一个子集
	item := make([]int, len(track))
	copy(item, track)
	*res = append(*res, item)

	// 回溯算法标准框架
	for i := index; i < len(nums); i++ {
		// 做选择
		track = append(track, nums[i])
		// 通过 start 参数控制树枝的遍历，避免产生重复的子集
		backtrack78(res, track, nums, i+1)
		// 撤销选择
		track = track[:len(track)-1]
	}
}
