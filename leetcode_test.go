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
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
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
