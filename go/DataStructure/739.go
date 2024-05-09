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
		stack = append(stack, i)
	}
	return res
}

// 从左往右寻找,依次加入栈,遇到比栈顶大的就出栈并计算距离
func dailyTemperatures2(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{}
	for i, temperature := range temperatures {
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			res[prevIndex] = i - prevIndex
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
