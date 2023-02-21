package DP

func canJump(nums []int) bool {
    n := len(nums)
    farthest := 0
    for i := 0; i < n; i++ {
        if i <= farthest {
            farthest = max(farthest, i+nums[i])
            if farthest >= n-1 {
                return true
            }
        }
    }
    return false
}
