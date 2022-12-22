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
