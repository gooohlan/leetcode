package BinarySearch

// 此处f(x)是一个单调递减的函数,所以l和r的调整要调换下顺序
func minEatingSpeed(piles []int, h int) int {
	l, r := 1, getMaxPile(piles) // 最小速度为1
	for l < r {
		mid := int(uint(l+r) >> 1)
		// mid := l + (r - l) / 2
		if f(piles, mid) <= h {
			// 需要让 f(x) 的返回值大一些
			r = mid
		} else {
			// 需要让 f(x) 的返回值小一些
			l = mid + 1
		}
	}
	return l
}

// 定义：速度为 x 时，需要 f(x) 小时吃完所有香蕉
func f(piles []int, x int) int {
	var hours int
	for _, pile := range piles {
		hours += pile / x
		if pile%x > 0 {
			hours++
		}
	}
	return hours
}

func getMaxPile(piles []int) int {
	max := 0
	for _, pile := range piles {
		if max < pile {
			max = pile
		}
	}
	return max
}
