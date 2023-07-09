package WeeklyContest

func maxNonDecreasingLength(nums1 []int, nums2 []int) int {
    n := len(nums1)
    nums := [2][]int{nums1, nums2}
    memo := make([][2]int, n)
    for i := range memo {
        memo[i] = [2]int{-1, -1}
    }
    var dfs func(int, int) int // i,j 以下标i结束的nums[j]结尾的最长非递归子数组长度
    
    dfs = func(i int, j int) int {
        if i == 0 {
            return 1
        }
        
        if memo[i][j] == -1 {
            memo[i][j] = 1
            if nums1[i-1] <= nums[j][i] {
                memo[i][j] = dfs(i-1, 0) + 1
            }
            if nums2[i-1] <= nums[j][i] {
                memo[i][j] = max(dfs(i-1, 1)+1, memo[i][j])
            }
        }
        
        return memo[i][j]
    }
    
    var res int
    for j := 0; j < 2; j++ {
        for i := 0; i < n; i++ {
            res = max(res, dfs(i, j))
        }
    }
    return res
}

func maxNonDecreasingLengthDP(nums1 []int, nums2 []int) int {
    n := len(nums1)
    nums := [2][]int{nums1, nums2}
    dp := make([][2]int, n)
    dp[0] = [2]int{1, 1}
    
    var res = 1
    for i := 1; i < n; i++ {
        dp[i] = [2]int{1, 1}
        for j := 0; j < 2; j++ {
            if nums1[i-1] <= nums[j][i] {
                dp[i][j] = dp[i-1][0] + 1
            }
            if nums2[i-1] <= nums[j][i] {
                dp[i][j] = max(dp[i-1][1]+1, dp[i][j])
            }
        }
        res = max(res, max(dp[i][0], dp[i][1]))
    }
    
    return res
}
