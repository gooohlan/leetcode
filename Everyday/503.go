package Everyday

func nextGreaterElements(nums []int) []int {
    ans := make([]int, len(nums))
    var stack []int
    n := len(nums)
    
    for i := 2*n - 1; i > 0; i-- {
        x := nums[i%n]
        
        for len(stack) > 0 && stack[len(stack)-1] <= x {
            // 左侧树更大,栈顶无效,移除
            stack = stack[:len(stack)-1]
        }
        if len(stack) == 0 { // 右侧没有更大的
            ans[i%n] = -1
        } else {
            ans[i%n] = stack[len(stack)-1]
        }
        stack = append(stack, x)
    }
    return ans
}
