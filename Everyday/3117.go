package Everyday

import (
    "fmt"
    "math"
)

func minimumValueSum(nums []int, andValues []int) int {
    const inf = math.MaxInt / 2
    n, m := len(nums), len(andValues)
    memo := map[string]int{} // i-j-and
    
    var dfs func(int, int, int) int
    dfs = func(i, j, and int) int {
        if n-i < m-j { // 剩余元素不足
            return inf
        }
        if j == m {
            if i == n {
                return 0
            }
            return inf // 数组中还有元素未分配
        }
        
        and &= nums[i]
        memoKey := fmt.Sprintf("%d-%d-%d", i, j, and)
        if res, ok := memo[memoKey]; ok {
            return res
        }
        res := dfs(i+1, j, and)  // 不选择nums[i]为尾数
        if and == andValues[j] { // nums[i]可以是尾数
            res = min(res, dfs(i+1, j+1, 1)+nums[i])
        }
        memo[memoKey] = res
        return res
    }
    ans := dfs(0, 0, -1) // -1&x = x
    if ans == inf {
        return -1
    }
    return ans
}
