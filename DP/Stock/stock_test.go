package Stock

import (
	"fmt"
	"testing"
)

func Test121(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxprofit121(prices))
	fmt.Println(maxProfit121_1(prices))
}

func Test122(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit122(prices))
	fmt.Println(maxProfit122_1(prices))
	fmt.Println(maxProfit122_2(prices))
}

func Test309(t *testing.T) {
	prices := []int{2, 1, 4}
	fmt.Println(maxProfit309(prices))
	fmt.Println(maxProfit309_1(prices))
}
