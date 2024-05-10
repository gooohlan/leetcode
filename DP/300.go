package DP

func lengthOfLIS(nums []int) int {
    dp := make([]int, len(nums))
    
    for i := 0; i < len(nums); i++ {
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                if dp[j]+1 > dp[i] {
                    dp[i] = dp[j] + 1
                }
            }
        }
    }
    
    res := 0
    for _, i := range dp {
        if i > res {
            res = i
        }
    }
    return res + 1
}

func lengthOfLIS2(nums []int) int {
    stack := []int{}
    for _, num := range nums {
        l := 0
        for r := len(stack); l < r; {
            mid := (l + r) / 2
            if stack[mid] >= num {
                r = mid
            } else {
                l = mid + 1
            }
        }
        if l == len(stack) {
            stack = append(stack, num)
        } else {
            stack[l] = num
        }
    }
    return len(stack)
}
