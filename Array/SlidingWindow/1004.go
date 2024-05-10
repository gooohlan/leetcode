package SlidingWindow

// 使用滑动窗口,当窗口中0的个数小于K时,我们向右扩大窗口
// 窗口中0个数大于k时向右缩小窗口
func longestOnes(nums []int, k int) int {
	l, r := 0, 0
	valid := 0
	res := 0

	for r < len(nums) {
		c := nums[r]
		r++

		if c == 0 {
			valid++
		}

		for valid > k {
			d := nums[l]
			l++
			if d == 0 {
				valid--
			}
		}

		res = max(r-l, res)
	}
	return res
}
