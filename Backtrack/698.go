package Backtrack

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
    n := len(nums)
    if k > n {
        return false
    }
    sum := 0
    for _, num := range nums {
        sum += num
    }
    if sum%k != 0 {
        return false
    }
    target := sum / k
    bucket := make([]int, k)
    
    var backtrack func(index int) bool
    backtrack = func(index int) bool {
        if index == n {
            for _, v := range bucket {
                if v != target {
                    return false
                }
            }
            return true
        }
        
        for i, v := range bucket {
            if v+nums[index] > target {
                continue
            }
            
            bucket[i] = v + nums[index]
            
            if backtrack(index + 1) {
                return true
            }
            bucket[i] -= nums[index]
        }
        
        return false
    }
    return backtrack(0)
}
func canPartitionKSubsets2(nums []int, k int) bool {
    n := len(nums)
    if k > n {
        return false
    }
    sum := 0
    for _, num := range nums {
        sum += num
    }
    if sum%k != 0 {
        return false
    }
    target := sum / k
    memo := make(map[int]bool) // 记录是否计算过
    used := 0                  // 使用位图计算是否使用过
    sort.Ints(nums)
    var backtrack func(k, bucket, start int) bool
    backtrack = func(k, bucket, start int) bool {
        if k == 0 {
            // 所有捅装满了
            // nums也用完了(sum/k==0)
            return true
        }
        
        if bucket == target {
            // 装满了当前桶，递归穷举下一个桶的选择
            // 让下一个桶从 nums[0] 开始选数字
            res := backtrack(k-1, 0, 0)
            memo[used] = res
            return res
        }
        
        if memo[used] {
            return memo[used]
        }
        
        for i := start; i < n; i++ {
            if used>>i&1 == 1 { // 判断第 i 位是否是 1
                // nums[i] 已经被装入别的桶中
                continue
            }
            if nums[i]+bucket > target {
                break
            }
            used |= 1 << i // 将第 i 位置为 1
            bucket += nums[i]
            
            if backtrack(k, bucket, i+1) {
                return true
            }
            used ^= 1 << i
            bucket -= nums[i]
        }
        
        return false
    }
    
    return backtrack(k, 0, 0)
}

// 官方解
func canPartitionKSubsets3(nums []int, k int) bool {
    all := 0
    for _, v := range nums {
        all += v
    }
    if all%k > 0 {
        return false
    }
    per := all / k
    sort.Ints(nums)
    n := len(nums)
    if nums[n-1] > per {
        return false
    }
    
    dp := make([]bool, 1<<n)
    for i := range dp {
        dp[i] = true
    }
    var dfs func(int, int) bool
    dfs = func(s, p int) bool {
        if s == 0 {
            return true
        }
        if !dp[s] {
            return dp[s]
        }
        dp[s] = false
        for i, num := range nums {
            if num+p > per {
                break
            }
            if s>>i&1 > 0 && dfs(s^1<<i, (p+nums[i])%per) {
                return true
            }
        }
        return false
    }
    return dfs(1<<n-1, 0)
}
