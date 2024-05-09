package WeeklyContest

// 在s中寻找t,有几个字符没找到就添加几个字符
func appendCharacters(s string, t string) int {
	i, j := 0, 0
	for ; i < len(s) && j < len(t); i++ {
		if s[i] == t[j] {
			j++
		}
	}
	return len(t) - j
}
