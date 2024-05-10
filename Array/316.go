package Array

// 整个题的要求,我们可以分为3点
// 去重
// 去重后字符串不能打乱原字符串出现的相对顺序
// 字典序最小的座位最终结果
// 我们采用栈的方式存储出现的字符,并有hash记录是否出现,这样遍历完一次栈中的数就是去重且保证出现了相对顺序的结果,满足了前两个要求
// 对于第三点, 假设我们栈中存在'b'和'c',插入'a'时,我们发现'a'比'b'和'c'都小,此时我们前面的字符pop出来,把'a'放在最前面
// 这样也有问题, 假设s = "bcac", 安装上面的方法得出的是'ac', 而正确答案是'bac',因为'b'在字符串中只出现了一次,所以我们出栈是需要判断,后续还会出现该字符时才出栈

// 用数组模拟一个本题需要的栈
type Stack struct {
	Arr []byte
}

func NewStack() *Stack {
	return &Stack{Arr: make([]byte, 0)}
}

func (s *Stack) Push(val byte) {
	s.Arr = append(s.Arr, val)
}

func (s *Stack) Pop() byte {
	n := s.Top()
	s.Arr = s.Arr[:len(s.Arr)-1]
	return n
}

func (s *Stack) isEmpty() bool {
	return len(s.Arr) == 0
}

func (s *Stack) Top() byte {
	return s.Arr[len(s.Arr)-1]
}

func removeDuplicateLetters(s string) string {
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
