package main

import "log"

func containsNearbyDuplicate(nums []int, k int) bool {
	mapArr := make(map[int]int)
	for key, val := range nums {
		if v, ok := mapArr[val]; ok {
			if key-v <= k {
				return true
			}
		}
		mapArr[val] = key
	}
	return false
}

func main()  {
	log.Println(containsNearbyDuplicate([]int{1,2,3,1,4}, 2))
}
