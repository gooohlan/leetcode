package WeeklyContest

import "math"

func pivotInteger(n int) int {
	// 构建前序和数组
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + i
	}

	for i := 1; i <= n; i++ {
		if sum[i] == sum[n]-sum[i-1] {
			return i
		}
	}

	return -1
}

// 数学
// [1,x]的和为:x(x+1)/2, [1,n]的和为n(n+1)/2, [x,n]的和为(n(n+1)-x(x-1))/2
// 因为[1,x]与[x,n]相等,简化后的等式为 x=√n(n+1)/2
// 此时如果x非整数则可以直接返回-1
func pivotIntegerMath(n int) int {
	n = n * (n + 1) / 2
	x := int(math.Sqrt(float64(n)))
	if x*x == n { // x 非整数无法等于n
		return x
	}
	return -1
}
