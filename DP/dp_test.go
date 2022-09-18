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
