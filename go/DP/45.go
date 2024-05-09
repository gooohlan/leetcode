package DP

func jump(nums []int) int {
    n := len(nums)
    end, farthest := 0, 0
    jumps := 0
    for i := 0; i < n-1; i++ {
        farthest = max(farthest, i+nums[i]) // 找到当前位置可达最远距离
        if i == end {                       // 到达上一次最为之,跳跃一次
            jumps++
            end = farthest // 更新最远位置
        }
    }
    return jumps
}
