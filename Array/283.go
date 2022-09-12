package Array

// 快指针fast向前遍历, 遇到非0值就与慢指针slow所在值进行交换,快指针遍历结束,所有的0都处于数组后半段
// 此题与26和27的思路类似,可以看做是27题的移除元素0的升级版,把移除变成了交换位置
func moveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
	}
}
