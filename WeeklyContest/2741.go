package WeeklyContest

func specialPerm(nums []int) int {
    const mod int = 10e9 + 7
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
