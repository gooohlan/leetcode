package main

import (
	"fmt"
	"testing"
)

func Test204(t *testing.T) {
	fmt.Println(countPrimes(10))
	fmt.Println(countPrimes2(10))
	fmt.Println(countPrimes3(10))
}

func Test303(t *testing.T) {
	obj := Constructor1([]int{-2, 0, 3, -5, 2, -1})
	param_1 := obj.SumRange(0, 5)
	fmt.Println(param_1)
	obj = Constructor2([]int{-2, 0, 3, -5, 2, -1})
	param_1 = obj.SumRange2(0, 5)
	fmt.Println(param_1)
}

func Test349(t *testing.T) {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersection2([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection2([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}

func Test350(t *testing.T) {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersect2([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect2([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}

func Test453(t *testing.T) {
	fmt.Println(minMoves([]int{5, 7, 1, 9}))
	fmt.Println(minMoves2([]int{5, 7, 1, 9}))
	fmt.Println(minMoves3([]int{3, 10, 15}))
	fmt.Println(minMoves3([]int{13, 18, 3, 10, 35, 68, 50, 20, 50}))
}

func Test455(t *testing.T) {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
}

func Test496(t *testing.T) {
	fmt.Println(nextGreaterElement2([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement2([]int{2, 4}, []int{1, 2, 3, 4}))
}

func Test506(t *testing.T) {
	fmt.Println(findRelativeRanks([]int{3, 5, 6, 1, 4}))
}

func Test598(t *testing.T) {
	ops := [][]int{
		[]int{2, 2},
		[]int{3, 3},
	}
	fmt.Println(maxCount(3, 3, ops))
}

func Test599(t *testing.T) {
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	fmt.Println(findRestaurant(list1, list2))
}

func Test645(t *testing.T) {
	fmt.Println(findErrorNums2([]int{1, 2, 2, 4}))
}
