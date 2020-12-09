package main

import (
	"fmt"
)

// 哈希
func findDisappearedNumbers(nums []int) []int {
	var rArr []int
	mapArr := make(map[int]bool)
	for _, v := range nums {
		mapArr[v] = true
	}
	for k, _ := range nums {
		if ok := mapArr[k+1]; !ok {
			rArr = append(rArr, k+1)
		}
	}
	return rArr
}

// 原地修改
func findDisappearedNumbers1(nums []int) []int {
	for _, v := range nums {
		if nums[abs(v)-1] > 0 {
			nums[abs(v)-1] *= -1
		}
	}
	var rArr []int
	for k,v := range nums{
		if v > 0 {
			rArr = append(rArr, k+1)
		}
	}
	return rArr
}

func abs(n int) int  {
	if n < 0 {
		n *=-1
	}
	return n
}

func main() {
	arr := findDisappearedNumbers1([]int{4, 3, 2, 7, 8, 2, 3, 1})
	fmt.Println(arr)
}
