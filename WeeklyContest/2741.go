package WeeklyContest

func specialPerm(nums []int) int {
    const mod int = 1e9 + 7
    n := len(nums)
    m := 1 << n
    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    
    var dfs func(int, int) int
    dfs = func(i int, j int) (res int) {
        if i == 0 {
            return 1
        }
        
        if memo[i][j] != -1 {
            return memo[i][j]
        }
        
        for k, x := range nums {
            if i>>k&1 == 1 && (nums[j]%x == 0 || x%nums[j] == 0) {
                res = (res + dfs(i^(1<<k), k)) % mod
            }
        }
        memo[i][j] = res
        return res
    }
    var res int
    for i := range nums {
        res = (res + dfs((m-1)^(1<<i), i)) % mod
    }
    return res
}

func specialPerm1(nums []int) int {
    const mod int = 1e9 + 7
    used := make(map[int]bool)
    var dfs func(track []int, j int) int
    
    dfs = func(track []int, j int) (res int) {
        if len(track) == len(nums) {
            return 1
        }
        
        for k, x := range nums {
            if !used[k] && (nums[j]%x != 0 || x%nums[j] != 0) {
                track = append(track, x)
                used[k] = true
                res = (res + dfs(track, k)) % mod
                used[k] = false
                track = track[:len(track)-1]
            }
        }
        return res
    }
    var res int
    for i, num := range nums {
        used[i] = true
        res = (res + dfs([]int{num}, i)) % mod
        used[i] = false
    }
    return res
}

func specialPermDP(nums []int) int {
    const mod int = 1e9 + 7
    n := len(nums)
    m := 1<<n - 1
    dp := make([][]int, m)
    dp[0] = make([]int, n)
    for i := range dp[0] {
        dp[0][i] = 1
    }
    
    for i := 1; i < m; i++ {
        dp[i] = make([]int, n)
        for j, x := range nums {
            for k, y := range nums {
                if i>>k&1 == 1 && (y%x == 0 || x%y == 0) {
                    dp[i][j] = (dp[i][j] + dp[i^(1<<k)][k]) % mod
                }
            }
        }
    }
    
    var res int
    for i := range nums {
        res = (res + dp[m^(1<<i)][i]) % mod
    }
    return res
}
