package DP

func findTargetSumWays(nums []int, target int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    
    var backtrack func(int, int)
    var res int
    backtrack = func(i, target int) {
        if i == n {
            if target == 0 {
                res++
            }
            return
        }
        
        // nums[i]赋值减号,目标就变成了target+nums[i]
        backtrack(i+1, target+nums[i])
        // nums[i]赋值加号,目标就变成了target-nums[i]
        backtrack(i+1, target-nums[i])
    }
    
    backtrack(0, target)
    return res
}

func findTargetSumWaysDFS(nums []int, target int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    memo := make(map[int]map[int]int, n)
    for i := 0; i < n; i++ {
        memo[i] = make(map[int]int)
    }
    
    var dfs func(int, int) int
    
    dfs = func(i, target int) int {
        if i == n {
            if target == 0 {
                return 1
            }
            return 0
        }
        
        if _, ok := memo[i][target]; !ok {
            memo[i][target] = dfs(i+1, target-nums[i]) + dfs(i+1, target+nums[i])
        }
        
        return memo[i][target]
    }
    
    return dfs(0, target)
}

// 把nums分为两个集合A和B,分别代表分配+和-的数,那么他们和target存在如下关系:
// sum(A) - sum(B) = target
// (sum - sum(B)) - sum(B) = target
// sum - target = 2 * sum(B)
// sum(B) = (sum - target) / 2
// 此时我们只需要求出和为sum(A)的即可
// 如果sum - target为负或者不能被2整除直接返回0
// 剩下的就又变成了背包问题,在nums数组中找出和为sum(A)的次数即可
func findTargetSumWaysDP(nums []int, target int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    diff := sum - target
    if diff < 0 || diff%2 != 0 {
        return 0
    }
    target = diff / 2
    n := len(nums)
    dp := make([][]int, n+1) // dp[i][j] 使用前i个元素凑出j的方法
    for i := range dp {
        dp[i] = make([]int, target+1)
    }
    dp[0][0] = 1 // 一个元素不使用凑出0的方法是1
    for i := 1; i <= n; i++ {
        for j := 0; j <= target; j++ {
            if j >= nums[i-1] {
                dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
            } else {
                dp[i][j] = dp[i-1][j]
            }
        }
    }
    return dp[n][target]
}

// 优化为以为
