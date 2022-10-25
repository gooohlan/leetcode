package BinarySearch

func shipWithinDays(weights []int, days int) int {
	l, r := 0, 0
	for _, weight := range weights {
		if weight > l {
			l = weight
		}
		r += weight
	}

	for l < r {
		mid := int(uint(l+r) >> 1)
		if f1101(weights, mid) <= days {
			// 需要让 f(x) 的返回值大一些
			r = mid
		} else {
			// 需要让 f(x) 的返回值小一些
			l = mid + 1
		}
	}
	return l
}

// 每天可运x,需要多少天
func f1101(weights []int, x int) int {
	days := 0

	for i := 0; i < len(weights); {
		cap := x
		for i < len(weights) {
			if cap < weights[i] {
				break
			}
			cap -= weights[i]
			i++
		}
		days++
	}
	return days
}
