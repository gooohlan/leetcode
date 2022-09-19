package Backtracking

var res [][]int

func permute(nums []int) [][]int {
	track := []int{}
	used := make(map[int]bool)
	backtrack(nums, track, used)
	return res
}

func backtrack(nums, track []int, used map[int]bool) {
	if len(nums) == len(track) {
		res = append(res, track)
		return
	}

	for i, num := range nums {
		if used[i] {
			continue
		}
		track = append(track, num)
		used[i] = true
		backtrack(nums, track, used)
		used[i] = false
		track = track[:len(track)-1]
	}
}

func permute1(nums []int) [][]int {
	var res2 [][]int
	backtrack2(&res2, 0, nums)
	return res2
}

func backtrack2(res2 *[][]int, first int, output []int) {
	if first == len(output) {
		list := make([]int, len(output))
		copy(list, output)
		*res2 = append(*res2, list)
	}
	for i := first; i < len(output); i++ {
		output[i], output[first] = output[first], output[i]
		backtrack2(res2, first+1, output)
		output[i], output[first] = output[first], output[i]
	}
}
