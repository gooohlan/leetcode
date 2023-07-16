package WeeklyContest

import "sort"

// 两人分队,所以最小的必须要和最大的匹配,否则和就不相同
func dividePlayers(skill []int) int64 {
	sort.Ints(skill)
	n := len(skill)
	sum := skill[0] + skill[n-1]
	var res int64
	for i := 0; i < len(skill)/2; i++ {
		a, b := skill[i], skill[n-1-i]
		if a+b != sum {
			return -1
		}
		res += int64(a * b)
	}
	return res
}
