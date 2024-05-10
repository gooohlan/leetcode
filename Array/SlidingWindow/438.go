package SlidingWindow

// 方法一对于方法二来说有所赘余,但是好处是这样可以称体系框架解决问题
func findAnagrams(s string, p string) []int {
	window, need := map[byte]int{}, map[byte]int{}
	for _, v := range p {
		need[byte(v)]++
	}

	l, r := 0, 0
	valid := 0
	res := []int{}

	for r < len(s) {
		c := s[r]
		r++

		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		if r-l >= len(p) {
			if valid == len(need) {
				res = append(res, l)
			}
			d := s[l]
			l++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}

func findAnagrams2(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}

	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return
}
