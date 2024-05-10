package Array

import "sort"

// 将nums1进行排序,nums2的索引进行排序
// 通过对比修改nums2, 最终返回nums2
func advantageCount(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	n := len(nums1)
	idx := make([]int, n)
	for i := 1; i < n; i++ {
		idx[i] = i
	}
	// 对nums2的索引进行
	sort.Slice(idx, func(i, j int) bool { return nums2[idx[i]] < nums2[idx[j]] })

	// nums1[left] 是最小值，nums1[right] 是最大值
	l, r := 0, n-1
	for _, num := range nums1 {
		// 如过nums1大过nums2,就安排到当前位置, 否则将这个拿去和nums2最大值进行比较
		if num > nums2[idx[l]] {
			nums2[idx[l]] = num
			l++
		} else {
			nums2[idx[r]] = num
			r--
		}
	}
	return nums2
}
