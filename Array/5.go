package Array

// 要判断一个字符串是否是回文数最好的方式就是从中心向两段扩散
// 回文串可以分为两种,长度为奇数或者偶数的
// 如果回文串长度为奇数,那就冲中心节点向两段扩散,是偶数则从中间两个节点向两段扩散
// 我们遍历给到我们的字符串中的每个字符,依次求出他们对应的回文串,最后取出最长的即可
func getPalindrome(s string, i, j int) string {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}

	return s[i+1 : j]
}

func longestPalindrome(s string) string {
	var str string
	for i := 0; i < len(s); i++ {
		res := getPalindrome(s, i, i)
		if len(res) > len(str) {
			str = res
		}
		res = getPalindrome(s, i, i+1)
		if len(res) > len(str) {
			str = res
		}
	}
	return str
}
