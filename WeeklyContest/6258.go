package WeeklyContest

import "sort"

func longestSquareStreak(nums []int) int {
	sort.Ints(nums)
	mapInt := make(map[int]int)
	for _, num := range nums {
		mapInt[num] = 1
	}
	var res = -1

	for _, num := range nums {
		if mapInt[num*num] != 0 {
			mapInt[num*num] = mapInt[num] + 1
		}
	}

	for _, i := range mapInt {
		if i > 1 {
			res = max(i, res)
		}
	}
	return res
}
