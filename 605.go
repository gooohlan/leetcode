package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 1}, 1))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed) && n > 0; {
		if flowerbed[i] == 1 { // 当前格子已经有花,我们两格后见
			i += 2
		} else if i == len(flowerbed)-1 || flowerbed[i+1] == 0 { // 当前格子不是花,下一个格子也不是花或着当前格子已经是最后一个
			n--
			i += 2
		} else { // 当前格子不是花,但是下一格是
			i += 3
		}
	}
	return n == 0
}
