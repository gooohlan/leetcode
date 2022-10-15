package Array

// 定义左右指针,相向而行,交换值
func reverseString(s []byte) []byte {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return s
}
