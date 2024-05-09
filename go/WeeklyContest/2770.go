package WeeklyContest

import "math"

func maximumJumps(nums []int, target int) int {
    n := len(nums)
    var dfs func(int) int
    memo := make([]int, n)
    for i := 0; i < n; i++ {
        memo[i] = -1
    }
    dfs = func(index int) int {
        if index == 0 {
            return 0
        }
        
        if memo[index] == -1 {
            memo[index] = math.MinInt
            for i := index - 1; i >= 0; i-- {
                if -target <= nums[index]-nums[i] && nums[index]-nums[i] <= target {
                    memo[index] = max(memo[index], dfs(i)+1)
                }
            }
            
        }
        return memo[index]
    }
    res := dfs(n - 1)
    if res < 0 {
        return -1
    }
    return res
}

// func abs(a, b int) int {
//     if a > b {
//         return a - b
//     }
//     return b - a
// }
