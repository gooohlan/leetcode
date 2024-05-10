package DP

import "math"

// 暴力解法
func coinChange(coins []int, amount int) int {
	return dp(coins, amount)
}

// 要凑出金额 n，至少要 dp(coins, n) 个硬币
func dp(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res := math.MaxInt
	for _, coin := range coins {
		// 计算子问题的结果
		number := dp(coins, amount-coin)
		// 子问题无解则跳过
		if number == -1 {
			continue
		}
		// 在子问题中选择最优解，然后加一
		if res > number {
			res = number
		}
	}
	if res == math.MaxInt {
		return -1
	}
	return res + 1
}

var memo []int

// 优化1:带备忘的递归
func coinChange1(coins []int, amount int) int {
	memo = make([]int, amount+1)
	for i, _ := range memo {
		// 备忘录初始化为一个不会被取到的特殊值，代表还未被计算
		memo[i] = -2
	}
	return dp1(coins, amount)
}

func dp1(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	// 查备忘录，防止重复计算
	if memo[amount] != -2 {
		return memo[amount]
	}

	res := math.MaxInt
	for _, coin := range coins {
		// 计算子问题的结果
		number := dp1(coins, amount-coin)
		// 子问题无解则跳过
		if number == -1 {
			continue
		}
		// 在子问题中选择最优解，然后加一
		if res > number {
			res = number
		}
	}

	// 把计算结果存入备忘录
	if res == math.MaxInt {
		memo[amount] = -1
		return -1
	} else {
		memo[amount] = res + 1
	}
	return memo[amount]
}

// 优化2:DP数组
func coinChange2(coins []int, amount int) int {
	// 数组大小为 amount + 1，初始值也为 amount + 1
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	// 外层 for 循环在遍历所有状态的所有取值
	for i := 0; i <= amount; i++ {
		// 内层 for 循环在求所有选择的最小值
		for _, coin := range coins {
			// 子问题无解，跳过
			if i-coin < 0 {
				continue
			}
			// dp[i-coin] == dp(coins, amount-coin)
			if dp[i] > dp[i-coin] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
