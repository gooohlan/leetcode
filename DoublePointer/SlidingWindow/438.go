package SlidingWindow

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

		for valid
	}
}
