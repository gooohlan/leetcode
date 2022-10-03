package SlidingWindow

import "math"

func minWindow(s string, t string) string {
	window, need := map[byte]int{}, map[byte]int{}
	for _, v := range t {
		need[byte(v)]++
	}

	l, r := 0, 0
	valid := 0                      // 记录符合条件的字符数
	start, length := 0, math.MaxInt // 开始位置和长度
	for r < len(s) {
		c := s[r]
		r++
		if _, ok := need[c]; ok {
			window[c]++
			if need[c] == window[c] {
				valid++
			}
		}

		// 判断左侧窗口是否要收缩
		for valid == len(need) {
			if r-l < length {
				length = r - l
				start = l
			}

			// d 是将移出窗口的字符
			d := s[l]
			// 窗口收缩
			l++
			if _, ok := need[d]; ok {
				if window[d] == need[d] { // 如果缩小窗口时移除后不符合需求字符,这退出sshou
					valid--
				}
				window[d]--
			}
		}
	}
	if length == math.MaxInt {
		return ""
	}
	return s[start : start+length]
}
