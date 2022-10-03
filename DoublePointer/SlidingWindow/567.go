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

		for r-l >= len(s1) {
			if valid == len(need) {
				return true
			}
			d := s2[l]
			l++
			// 在这里判断是否找到了合法的子串
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return false
}
