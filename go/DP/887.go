package DP

import (
    "math"
)

func superEggDrop(k int, n int) int {
    var dfs func(int, int) int
    memo := make([][]int, k+1)
    for i := range memo {
        memo[i] = make([]int, n+1)
    }
    dfs = func(K int, N int) int {
        if K == 1 {
            return N
        }
        if N == 0 {
            return 0
        }
        if memo[K][N] == 0 {
            res := math.MaxInt
            
            for i := 1; i <= N; i++ {
                res = min(res, max(
                    dfs(K, N-i),   // 没碎
                    dfs(K-1, i-1), // 碎了
                )+1)
            }
            memo[K][N] = res
        }
        
        return memo[K][N]
    }
    
    return dfs(k, n)
}

// 二分优化
func superEggDrop2(k int, n int) int {
    var dfs func(int, int) int
    memo := make([][]int, k+1)
    for i := range memo {
        memo[i] = make([]int, n+1)
    }
    // 代表着拥有K个鸡蛋，有N层楼需要遍历时，最坏情况下需要的最少次数
    dfs = func(K int, N int) int {
        // base case
        // 只有1个鸡蛋的时候，只能进行线性扫描法逐楼层遍历
        if K == 1 {
            return N
        }
        // 有0个楼层需要遍历，自然需要遍历0次
        if N == 0 {
            return 0
        }
        
        if memo[K][N] == 0 {
            res := math.MaxInt
            lo, hi := 1, N
            for lo <= hi {
                mid := (lo + hi) / 2
                broken := dfs(K-1, mid-1)  // 碎了
                notBroken := dfs(K, N-mid) // 没碎
                // 如果蛋碎了的次数比较大，则搜索范围为[1,mid-1]
                if broken > notBroken {
                    hi = mid - 1
                    res = min(res, broken+1)
                } else { // 如果蛋没碎的次数比较大，则搜索范围为[mid+1,N]
                    lo = mid + 1
                    res = min(res, notBroken+1)
                }
            }
            memo[K][N] = res
        }
        
        return memo[K][N]
    }
    
    return dfs(k, n)
}

// dp[i][j] = n
// 当前有 i 个鸡蛋，可以尝试扔 j 次鸡蛋
// 这个状态下，最坏情况下最多能确切测试一栋 n 层的楼
// 1、无论你在哪层楼扔鸡蛋，鸡蛋只可能摔碎或者没摔碎，碎了的话就测楼下，没碎的话就测楼上。
// 2、无论你上楼还是下楼，总的楼层数 = 楼上的楼层数 + 楼下的楼层数 + 1（当前这层楼）。
func superEggDropDP(k, n int) int {
    dp := make([][]int, k+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    
    j := 0
    for dp[k][j] < n {
        j++
        for i := 1; i <= k; i++ {
            // dp[i][j-1]   表示楼上的层数,因为i不变,鸡蛋没碎,但是扔鸡蛋次数j-1
            // dp[i-1][j-1] 表示楼上的层数,因为i-1,鸡蛋碎了,扔鸡蛋次数j-1
            // 1表示本层
            dp[i][j] = dp[i][j-1] + dp[i-1][j-1] + 1
        }
    }
    return j
}

func superEggDropDP2(k, n int) int {
    dp := make([]int, k+1)
    
    j := 0
    for dp[k] < n {
        j++
        pre := 0
        dp[0] = 0
        for i := 1; i <= k; i++ {
            // dp[i][j-1]   表示楼上的层数,因为i不变,鸡蛋没碎,但是扔鸡蛋次数j-1
            // dp[i-1][j-1] 表示楼上的层数,因为i-1,鸡蛋碎了,扔鸡蛋次数j-1
            // 1表示本层
            tmp := dp[i]
            dp[i] = dp[i] + pre + 1
            // dp[i][j] = dp[i][j-1] + dp[i-1][j-1] + 1
            pre = tmp
        }
    }
    return j
}
