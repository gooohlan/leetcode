package SlidingWindow

import "fmt"

func strStr(haystack string, needle string) int {
	i, j := 0, 0
	for ; i < len(haystack) && j < len(needle); i++ {
		if haystack[i] == needle[j] { // 相等则继续向后匹配
			j++
		} else { // 不相等指针会带原处
			i = i - j
			j = 0
		}
	}
	if j == len(needle) {
		return i - j
	}
	return -1
}

// Rabin-Karp 解法
func strStr2(haystack string, needle string) int {

	// 进制（只考虑 ASCII 编码）
	var R = 256
	// 数字位数
	var L = len(needle)
	// 取一个比较大的素数作为求模的除数
	var Q = 1658598167
	// R^(L - 1) 的结果
	var RL = 1
	for i := 1; i <= L-1; i++ {
		// 计算过程中不断求模，避免溢出
		RL = (RL * R) % Q
	}

	// 计算模式串的哈希值，时间 O(L)
	var patHash = 0
	for i := 0; i < len(needle); i++ {
		patHash = (R*patHash + int(needle[i])) % Q
	}
	x := 0
	// 存储前9个数字
	for i := 0; i < L-1; i++ {
		x = x*R%Q + int(haystack[i])%Q
	}

	for i := 0; i <= len(haystack)-L; i++ {
		x = x*R%Q + int(haystack[i+L-1])%Q // 从右边添加字符串

		if x == patHash {
			fmt.Println(haystack[i : i+L])
			if haystack[i:i+L] == needle {
				return i
			}
		}
		x = (x - int(haystack[i])*RL%Q + Q) % Q
	}
	return -1
}
