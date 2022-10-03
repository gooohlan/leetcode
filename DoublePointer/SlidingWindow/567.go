package SlidingWindow

// 简单版本的76
func checkInclusion(s1 string, s2 string) bool {
	window, need := map[byte]int{}, map[byte]int{}
	for _, v := range s1 {
		need[byte(v)]++
	}

	l, r := 0, 0
	valid := 0

	for r < len(s2) {
		c := s2[r]
		r++
		if _, ok := need[c]; ok {
			window[c]++
			if need[c] == window[c] {
				valid++
			}
		}

		// 因为是排列,所以左侧收缩的条件为整个字符串的长度大于子串长度
		if r-l >= len(s1) {
			if valid == len(need) {
				return true
			}
			d := s2[l]
			l++
			// 在这里判断是否找到了合法的子串
			if _, ok := need[d]; ok {
				window[d]--
				if window[d] == need[d] {
					valid--
				}
			}
		}
	}

	return false
}
