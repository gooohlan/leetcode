package Merge

// 归并排序, 取一个中间值,先排左边,再排右边,然后合并
func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])

	res := make([]int, len(left)+len(right))
	l, r := 0, 0
	for i := range nums {
		if l < len(left) && (r == len(right) || left[l] <= right[r]) {
			res[i] = left[l]
			l++
		} else {
			res[i] = right[r]
			r++
		}
	}

	return res
}
