package Everyday

func minDays(n int) int {
    memo := map[int]int{}
    var dfs func(int) int

    dfs = func(i int) int {
        if i <= 1 {
            return 1
        }

        if _, ok := memo[i]; !ok {
            memo[i] = 1 + min(dfs(i/2)+i%2, dfs(i/3)+i%3)
        }

        return memo[i]
    }

    return dfs(n)
}
