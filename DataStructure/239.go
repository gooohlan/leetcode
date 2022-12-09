package DataStructure

// 使用单调队列,维护一个长度为K的队列,队列内保证是单调递增的,
// 每次入队维护当前队列,然后找出队列中最大的,如果此时找出的最大值不在范围内,则移除队首,继续寻找下一个最大值
func maxSlidingWindow(nums []int, k int) []int {
	var (
		q    []int // 存储数组下标,方便后续判断是否在范围内
		push func(int)
		res  []int
	)

	push = func(x int) {
		for len(q) > 0 && nums[q[len(q)-1]] < nums[x] { // 如果队尾小于插入的元素,需要移除
			q = q[:len(q)-1] // 移除最后一个元素
		}
		q = append(q, x)
	}

	for i := 0; i < k; i++ {
		push(i)
	}
	res = append(res, nums[q[0]])

	for i := k; i < len(nums); i++ {
		push(i)
		for q[0] <= i-k { // 最大值不在k有效范围内,移除
			q = q[1:]
		}
		res = append(res, nums[q[0]])
	}

	return res
}
