package main

func convertToTitle(n int) string {
	var str string
	for n > 0 {
		n-- // 减去一个,因为A是对应的是1,而不是0
		str = string('A'+int32(n%26)) + str
		n /= 26
	}
	return str
}
