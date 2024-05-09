package Array

func reverseWords(s string) string {
	n := len(s)
	res := []byte{}
	for i := 0; i < n; {
		left := i
		for i < n && s[i] != ' ' {
			i++
		}
		for right := i - 1; right >= left; right-- {
			res = append(res, s[right])
		}

		for i < n {
			res = append(res, ' ')
			i++
		}
	}
	return string(res)
}
