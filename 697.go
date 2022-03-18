package main

type entry struct {
	cnt, l, r int
}

func findShortestSubArray(nums []int) int {
	var ans, maxCnt int
	mp := map[int]entry{}
	for i, v := range nums {
		if e, ok := mp[v]; ok {
			e.cnt++
			e.r = i
			mp[v] = e
		} else {
			mp[v] = entry{1, i, i}
		}
	}
	for _, e := range mp {
		if e.cnt > maxCnt {
			maxCnt, ans = e.cnt, e.r-e.l+1
		}
		if e.cnt == maxCnt {
			ans = min(ans, e.r-e.l+1)
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
