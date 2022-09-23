package Backtrack

import "sort"

//
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
