package DataStructure

// 将数组长度扩展为2倍长,即可获得所有正确答案
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}
	var stack []int
	for i := 0; i < n*2-1; i++ {
		for len(stack) > 0 && nums[i%n] > nums[stack[len(stack)-1]%n] {
			prevIndex := stack[len(stack)-1]
			res[prevIndex%n] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return res
}
