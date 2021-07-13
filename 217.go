package main

func containsDuplicate(nums []int) bool {
	mapArr := make(map[int]bool)
	for _, v := range nums {
		if mapArr[v] {
			return true
		}
		mapArr[v] = true
	}
	return false
}
