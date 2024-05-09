package Array

// 与316一样
func smallestSubsequence(s string) string {
	stack := NewStack()
	// 记录字符出现次数
	count := make([]byte, 256)
	for i := range s {
		count[s[i]]++
	}

	// 记录栈中是否存在
	inStack := make(map[byte]bool)

	for i := range s {
		count[s[i]]--

		// 已经存在栈中直接跳过
		if inStack[s[i]] {
			continue
		}

		// 如果栈不为空,且当前入栈字符串小于栈顶字符串,且字符串后续还存在该字符,则出栈
		for !stack.isEmpty() && stack.Top() > s[i] && count[stack.Top()] > 0 {
			inStack[stack.Pop()] = false
		}
		inStack[s[i]] = true
		stack.Push(s[i])
	}

	return string(stack.Arr)
}
