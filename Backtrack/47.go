package Backtrack

import (
	"math"
	"sort"
)

// 对于元素可重不可复选的,首先进行排序, 然后再减枝
// 剪枝一般处理为判断相邻两个数值是否一样
// 这里有个小技巧, 比如数组[1, 2, 2'] 与 [1, 2', 2]应该被算作同一个排列
// 要保证他们的不重复,就是要保证他的相对位置, 即保证 2 -> 2', 这样就能得到不重复的排列结果
// 当出现重复元素时，比如输入 nums = [1,2,2',2'']，2' 只有在 2 已经被使用的情况下才会被选择，同理，2'' 只有在 2' 已经被使用的情况下才会被选择，这就保证了相同元素在排列中的相对位置保证固定。
func permuteUnique(nums []int) [][]int {
	var (
		res       [][]int
		track     []int
		backtrack func()
	)
	used := make(map[int]bool, 0)

	sort.Ints(nums)
	backtrack = func() {
		if len(track) == len(nums) {
			item := make([]int, len(track))
			copy(item, track)
			res = append(res, item)
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			track = append(track, nums[i])
			used[i] = true
			backtrack()
			used[i] = false
			track = track[:len(track)-1]
		}
	}

	backtrack()
	return res
}

// 提供一个更简单,我们用一个值来存储已经遍历过的值,遇到相同的就跳过
// 比如 [1, 1, 3],遍历时遇到第二个1的时候就直接跳过
func permuteUnique2(nums []int) [][]int {
	var (
		res       [][]int
		track     []int
		backtrack func()
	)
	used := make(map[int]bool, 0)

	sort.Ints(nums)
	backtrack = func() {
		if len(track) == len(nums) {
			item := make([]int, len(track))
			copy(item, track)
			res = append(res, item)
		}

		// 上一个数字,用于标记是否重复
		prevNum := math.MaxInt
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			if nums[i] == prevNum {
				continue
			}

			prevNum = nums[i]
			track = append(track, nums[i])
			used[i] = true
			backtrack()
			used[i] = false
			track = track[:len(track)-1]
		}
	}

	backtrack()
	return res
}
