package WeeklyContest

// 如果 n <= 2,那么直接返回结果
// 当 n > 3 时无论怎么拆分,总会在某一时刻分割出一个长度为 2 的数组
// 如果 nums 所有长度为 2 的子数组和都小于 m, 那么肯定是false
func canSplitArray(nums []int, m int) bool {
    if len(nums) <= 2 {
        return true
    }
    
    n := len(nums)
    for i := 1; i < n; i++ {
        if nums[i-1]+nums[i] >= m {
            return true
        }
    }
    return false
}
