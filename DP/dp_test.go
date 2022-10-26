package DP

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	fmt.Println(fib1(20))
	fmt.Println(fib2(20))
	fmt.Println(fib3(20))
	fmt.Println(fib4(20))
}

func TestCoinChange(t *testing.T) {
	cois := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(cois, amount))
	fmt.Println(coinChange1(cois, amount))
	fmt.Println(coinChange2(cois, amount))
}

func Test198(t *testing.T) {
	nums := []int{1, 5, 2, 4, 1, 1, 6}
	fmt.Println(rob(nums))
	fmt.Println(rob1(nums))
	fmt.Println(rob2(nums))
	fmt.Println(rob3(nums))
}

func Test213(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob213(nums))
}

func Test337(t *testing.T) {
	node := &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, &TreeNode{Val: 5}}},
		Right: &TreeNode{9, nil, nil},
	}
	fmt.Println(rob337(node))
}
func Test410(t *testing.T) {
	nums := []int{7, 2, 5, 10, 8}
	fmt.Println(splitArrayDP(nums, 2))
}
