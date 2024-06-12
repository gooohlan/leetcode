package Everyday

func maxCoins(nums []int) int {
    n := len(nums)
    points := make([]int, n+2) // 收尾各插入一个1,方便操作
    points[0], points[n+1] = 1, 1
    copy(points[1:n+1], nums)
    
    dp := make([][]int, n+2)
    for i := range dp {
        dp[i] = make([]int, n+2)
    }
    
    for i := n; i >= 0; i-- {
        for j := i + 1; j < n+2; j++ {
            for k := i + 1; k < j; k++ {
                dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+points[i]*points[j]*points[k])
            }
        }
    }
    
    return dp[0][n+1]
}
