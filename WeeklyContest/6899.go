package WeeklyContest

import "math"

func maximumJumps(nums []int, target int) int {
    n := len(nums)
    var dfs func(int) int
    dp := make([]int, n)
    for i := 0; i < n; i++ {
        dp[i] = math.MaxInt
    }
    dfs = func(index int) int {
        if index == 0 {
            return 0
        }
        
        if dp[index] == math.MaxInt {
            dp[index] = -1
            for i := index - 1; i >= 0; i-- {
                if -target <= nums[index]-nums[i] && nums[index]-nums[i] <= target {
                    if dfs(i) != -1 {
                        dp[index] = max(dp[index], dfs(i)+1)
                    }
                }
            }
            
        }
        return dp[index]
    }
    return dfs(n - 1)
}

// func abs(a, b int) int {
//     if a > b {
//         return a - b
//     }
//     return b - a
// }
