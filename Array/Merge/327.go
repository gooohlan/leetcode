package Merge

// 先构件前缀和数组,归并是进行处理
// 给定两个升序排列的数组left, rifht分别找出下标1,j满足lower <= right[j]-[i] <= upper
// 在right种维护两个指针l,r都指向right初始位置
// l右移,直到left[0]+right[l]>=lower;r右移,直到left[0]+right[r]<upper
// 此时在[l,r)这个区间种所有的下标j,都满足lower <= right[j]-left[0] <= upper
// 因为left也是有序数组, 所以left后移时,l和r跟着后移就可以了
func countRangeSum(nums []int, lower int, upper int) int {

	var mergeCount func([]int) int
	mergeCount = func(arr []int) int {
		if len(arr) < 2 {
			return 0
		}

		mid := len(arr) / 2
		left := append([]int{}, arr[:mid]...)
		right := append([]int{}, arr[mid:]...)
		count := mergeCount(left) + mergeCount(right)

		l, r := 0, 0
		for _, v := range left {
			for l < len(right) && right[l]-v < lower { // left[0]+right[l]>=lower
				l++
			}

			for r < len(right) && right[r]-v <= upper { // left[0]+right[r]<upper
				r++
			}
			count += r - l
		}

		l, r = 0, 0
		for i := range arr {
			if l < len(left) && (r == len(right) || left[l] <= right[r]) {
				arr[i] = left[l]
				l++
			} else {
				arr[i] = right[r]
				r++
			}
		}
		return count
	}

	prefixSum := make([]int, len(nums)+1)
	for i, v := range nums {
		prefixSum[i+1] = prefixSum[i] + v
	}

	return mergeCount(prefixSum)
}
