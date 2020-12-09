package main

import (
	"fmt"
)

// 传统
func fib(N int) int {
	if N <= 1 {
		return N
	}
	f1, f2 := 1, 0
	for i := 2; i <= N; i++ {
		f1, f2 = f1+f2, f1
	}
	return f1
}

// 递归
func fib2(N int) int {
	if N <= 1 {
		return N
	}
	return fib(N-1) + fib(N-2)
}

// 带记忆的递归
func fib3(N int) int {
	if N <= 1 {
		return N
	}
	return memory(N)
}

var cache = map[int]int{0: 0, 1: 1}

func memory(N int) int {
	if v, ok := cache[N]; ok {
		return v
	}
	cache[N] = memory(N-1) + memory(N-2)
	return cache[N]
}

func main() {
	fmt.Println(fib2(4))
}
