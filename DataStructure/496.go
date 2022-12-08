package DataStructure

// 通过栈求出nums2中右侧比他大的元素,然后遍历num1获取对应值右侧比他大的元素
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack []int
	nextMap := make(map[int]int, len(nums2))
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && num >= stack[len(stack)-1] { // 当前元素大于栈顶
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 { // 栈不为空,右侧存在比自己大的元素
			nextMap[num] = stack[len(stack)-1]
		} else {
			nextMap[num] = -1
		}
		stack = append(stack, num)
	}

	res := make([]int, len(nums1))
	for i, num := range nums1 {
		res[i] = nextMap[num]
	}
	return res
}
