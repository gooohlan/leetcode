package SlidingWindow

// 相对于前几题,这个更简单了维护一个滑动窗口岂可
// 当 window[c]值大于1时，说明窗口中存在重复字符，不符合条件，就该移动left缩小窗口
func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	var res int
	window := map[string]int{}
	for r < len(s) {
		c := string(s[r])
		r++
		window[c]++
		// 去除窗口中所有的重复字符
		for window[c] > 1 {
			d := string(s[l])
			l++
			window[d]--
		}
		res = max(res, r-l)
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
