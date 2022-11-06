package Merge

// 依旧是有归并排序的思路作为主框架
// 根据题意,我们再合并时判断一下题目条件即可
// 对于每个左边的元素,我们都去右边寻找符合条件的元素
func reversePairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	mid := len(nums) / 2
	left := append([]int{}, nums[:mid]...)
	right := append([]int{}, nums[mid:]...)
	count := reversePairs(left) + reversePairs(right)

	// 合并时传递进来的数组都是排好序的,nums[i] <= nums[i+1]
	// 所以 当我们找到nums[i] > 2*nums[j]的nums[j],也必然也符合nums[i+1] > 2*nums[j]
	j := 0 // j用来记录右数组符合添加个数
	for i := 0; i < len(left); i++ {
		for j < len(right) && left[i] > right[j]*2 { // 符合条件时右侧数组指针后移
			j++
		}
		count += j
	}

	l, r := 0, 0
	for i := range nums {
		if l < len(left) && (r == len(right) || left[l] <= right[r]) {
			nums[i] = left[l]
			l++
		} else {
			nums[i] = right[r]
			r++
		}
	}

	return count
}
