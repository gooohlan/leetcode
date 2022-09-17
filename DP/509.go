package DP

// 暴力递归解法
func fib1(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib1(n-1) + fib1(n-2)
}

// 优化1:
// 带备忘录的递归解法
// 时间复杂度的优化,用空间换了时间
func fib2(n int) int {
	memo := make([]int, n+1)
	return helper(n, memo)
}

func helper(n int, memo []int) int {
	if n == 0 || n == 1 {
		return n
	}
	if memo[n] != 0 {
		// 已经计算过了
		return memo[n]
	}
	memo[n] = helper(n-1, memo) + helper(n-2, memo)
	return memo[n]
}

// 优化2:
// 使用dp数组完成自底向上的迭代
// dp数组与备忘录数组里存储的内容是一样的,把递归换成了for循环
func fib3(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	// 初始化 base case
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		// 状态转移
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 优化3:
// 每次计算的时候,其实我们只用到了前两个数组下标内存储的东西,存在则比较大的空间浪费
// 此时的算法时间复杂度为0(n),空间复杂度0(1)
func fib4(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		// 状态转移
		prev, curr = curr, prev+curr
	}
	return curr
}
