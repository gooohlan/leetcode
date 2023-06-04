package WeeklyContest

import "strconv"

func count(num1 string, num2 string, minSum int, maxSum int) int {
    MOD := int(1e9) + 7
    n1, _ := strconv.Atoi(num1)
    n2, _ := strconv.Atoi(num2)
    
    dp := make([][]int, len(num2)+1)
    for i := range dp {
        dp[i] = make([]int, maxSum+1)
    }
    
    dp[0][0] = 1
    
    for i := 1; i <= len(num2); i++ {
        for j := 0; j <= maxSum; j++ {
            for digit := 0; digit <= 9; digit++ {
                if n1 <= n2 && j >= digit {
                    newSum := j - digit
                    if newSum >= 0 && newSum <= maxSum {
                        dp[i][j] += dp[i-1][newSum]
                        dp[i][j] %= MOD
                    }
                }
            }
        }
    }
    
    result := 0
    for j := minSum; j <= maxSum; j++ {
        result += dp[len(num2)][j]
        result %= MOD
    }
    
    return result
}
