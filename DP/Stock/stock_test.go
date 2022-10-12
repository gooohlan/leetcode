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

func Test714(t *testing.T) {
	prices := []int{1, 3, 7, 5, 10, 3}
	fmt.Println(maxProfit714(prices, 3))
	fmt.Println(maxProfit714_1(prices, 3))
}

func Test123(t *testing.T) {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit123(prices))
	fmt.Println(maxProfit123_1(prices))
}

func Test188(t *testing.T) {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit188(2, prices))
}
