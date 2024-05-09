package main

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	var (
		res      = []string{}
		mapList  = make(map[string]int)
		minIndex = math.MaxInt32
	)
	for i, s := range list1 {
		mapList[s] = i
	}
	for i, s := range list2 {
		if v, ok := mapList[s]; ok {
			sum := v + i
			if sum == minIndex {
				res = append(res, s)
			}
			if sum < minIndex {
				minIndex = sum
				res = []string{s}

			}
		}
	}
	return res
}

