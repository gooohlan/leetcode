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

func Test303(t *testing.T){
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	param_1 := obj.SumRange(0, 5)
	fmt.Println(param_1)
	obj = Constructor2([]int{-2, 0, 3, -5, 2, -1})
	param_1 = obj.SumRange2(0, 5)
	fmt.Println(param_1)
}
