package Array

// 依旧采用快慢指针的方式
// 如果fast遇到值为val的元素,则直接跳过,否则赋值给slow所在值,slow前进一步
func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	// 因为先复制后前进,所以返回值不+1
	return slow
}
