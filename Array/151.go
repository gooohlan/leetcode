package Array

import "strings"

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	bytes := reverse([]byte(s), 0, len(s)-1)
	for i := 0; i < len(bytes); {
		for j := i; j < len(bytes); j++ {
			if j+1 == len(bytes) || bytes[j+1] == 32 {
				bytes = reverse(bytes, i, j)
				i = j + 2
				break
			}
		}
	}
	return string(bytes)
}

// 反转字符串并去除字符串之间多余空格
func reverse(s []byte, i, j int) []byte {
	for i < j {
		// 去除中间多余空格
		for s[i] == 32 && s[i+1] == 32 {
			s = append(s[:i], s[i+1:]...)
			j--
		}
		for j < len(s)-1 && s[j] == 32 && s[j-1] == 32 {
			s = append(s[:j], s[j+1:]...)
			j--
		}
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return s
}
