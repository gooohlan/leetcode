package Everyday

func specialPerm(nums []int) int {
    n := len(nums)
    memo := make([][]int, 1<<n-1)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1 // -1 表示没有计算过
        }
    }
    
    var dfs func(int, int) int
    dfs = func(s int, i int) int {
        if s == 1<<n {
            return 1 // 找到一个特别排列
        }
        
        if memo[s][i] != -1 {
            return memo[s][i]
        }
        
        res := 0
        for j, x := range nums {
            if s>>j&1 == 0 && (nums[i]%x == 0 || x%nums[i] == 0) {
                res += dfs(s|1<<j, j)
            }
        }
        
        memo[s][i] = res
        return res
    }
    
    ans := 0
    for i := range nums {
        ans += dfs(1<<i, i)
    }
    return ans % 1e9
}
