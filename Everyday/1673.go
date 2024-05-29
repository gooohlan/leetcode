package Everyday

func mostCompetitive(nums []int, k int) []int {
    stack := make([]int, 0)
    for i, num := range nums {
        for len(stack) > 0 && num < stack[len(stack)-1] && len(stack)+len(nums)-i > k {
            stack = stack[:len(stack)-1]
        }
        if len(stack) < k {
            stack = append(stack, num)
        }
    }
    return stack
}
