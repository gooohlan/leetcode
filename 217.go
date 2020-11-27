package main

import "log"

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

func main() {
	log.Println(containsDuplicate([]int{1, 2, 3, 4, 5, 3}))
}
