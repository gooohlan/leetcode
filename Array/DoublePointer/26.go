package Array

// 定义快慢指针,快指针fast在前面探路,慢指针slow跟在后面,找到一个不重复的元素就让慢指针slow前进一步并赋值
func removeDuplicates(nums []int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
