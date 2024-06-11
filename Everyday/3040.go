package Everyday

func maxOperations3040(nums []int) int {
    n := len(nums)
    res1, done := helper(nums[2:], nums[0]+nums[1])
    if done {
        return n / 2
    }
    res2, done := helper(nums[:n-2], nums[n-1]+nums[n-2])
    if done {
        return n / 2
    }
    res3, done := helper(nums[1:n-1], nums[0]+nums[n-1])
    if done {
        return n / 2
    }
    return max(res1, res2, res3) + 1
}

func helper(a []int, target int) (int, bool) {
    n := len(a)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    
    var done bool
    var dfs func(int, int) int
    dfs = func(i, j int) int {
        if done {
            return 0
        }
        if i >= j {
            done = true
            return 0
        }
        
        res := 0
        
        if memo[i][j] != -1 {
            return memo[i][j]
        }
        
        if a[i]+a[i+1] == target {
            res = max(res, dfs(i+2, j)+1)
        }
        if a[j]+a[j-1] == target {
            res = max(res, dfs(i, j-2)+1)
        }
        if a[i]+a[j] == target {
            res = max(res, dfs(i+1, j-1)+1)
        }
        memo[i][j] = res
        return res
    }
    
    return dfs(0, n-1), done
}
