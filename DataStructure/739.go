package DataStructure

// 使用单调栈,存储下标,遇到比组合大的减去下标即可
func dailyTemperatures(temperatures []int) []int {
    var stack []int
    res := make([]int, len(temperatures))
    for i := len(temperatures) - 1; i >= 0; i-- {
        for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
            stack = stack[:len(stack)-1]
        }
        
        if len(stack) > 0 {
            res[i] = stack[len(stack)-1] - i
        } else {
            res[i] = 0
        }
    }
    return res
}
