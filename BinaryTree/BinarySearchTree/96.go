package BinarySearchTree

func numTrees(n int) int {
	// 定义：闭区间 [lo, hi] 的数字能组成 count(lo, hi) 种 BST
	var count func(int, int) int
	memo := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, n+1)
	}

	count = func(lo, hi int) int {
		// 当 lo>hi时,闭区间[lo, hi]肯定是个空区间,即为nil,但是此时不能返回0,需要返回1
		if lo > hi {
			return 1
		}
		if memo[lo][hi] != 0 {
			return memo[lo][hi]
		}

		var res int
		for i := lo; i <= hi; i++ {
			left := count(lo, i-1)
			right := count(i+1, hi)
			res += left * right
		}
		memo[lo][hi] = res
		return res
	}

	return count(1, n)
}

// 上述代码可换成动态规划
func numTreesDP(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
