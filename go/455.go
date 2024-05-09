package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var count int
	for	i,j := 0,0; i<len(g) && j < len(s); {
		if g[i] <= s[j] { // 满足即分配
			count ++
			i++
			j++
		}else{ // 寻找符合孩子胃口的饼干
			j ++
		}
	}
	return count
}