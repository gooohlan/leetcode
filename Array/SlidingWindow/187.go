package SlidingWindow

import "math"

var codeMap = map[byte]int{
	'A': 0,
	'C': 1,
	'G': 2,
	'T': 3,
}

// 首先我们会议一下10进制数的删除最高位和添加最低位
// 添加最低位 432 *10 + 1 -> 4321
// 删除最高位 4321 - 4*10*3 -> 321 (3代表了位数)
// 假设我们是一个四进制数,那么他们的删除最高位和添加最低位就变成
// 添加最低位 432 *4 + 1 -> 4321
// 删除最高位 4321 - 4*4*3 -> 321
// 转换成代码,定义R为进制数, L为位数, removeVal 为最高位数字,appendVal为需要加入的数字
// 添加最低位 x = x * R + appendVal
// 删除最高位 x = x - removeVal*4*(L-1)
func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return []string{}
	}
	var res []string
	x := 0
	// 数字位数
	var L = 10
	// 进制
	var R = 4
	// 存储 R^(L - 1) 的结果
	var RL = int(math.Pow(float64(R), float64(L-1)))
	// 存储前9个数字
	for i := 0; i < 9; i++ {
		x = x*R + codeMap[s[i]]
	}

	cnt := map[int]int{}
	for i := 0; i <= len(s)-L; i++ {
		x = x*R + codeMap[s[i+L-1]] // 从右边添加字符串
		cnt[x]++                    // 记录这10位字符串出现次数
		if cnt[x] == 2 {
			res = append(res, s[i:i+L])
		}

		x = x - codeMap[s[i]]*RL
	}
	return res
}

// 以上面的思路,我们使用位运算来操作
// 添加最低位 x = x << 2 | appendVal
// 删除最高位 x = x & ((1 << 20) - 1) 由于我们只考虑 x 的低 20 位比特，需要将其余位置零，即与上 (1 << 20) - 1
func findRepeatedDnaSequences2(s string) []string {
	if len(s) < 10 {
		return []string{}
	}
	var res []string
	x := 0
	// 数字位数
	var L = 10
	// 存储前9个数字
	for i := 0; i < 9; i++ {
		x = x<<2 | codeMap[s[i]]
	}

	cnt := map[int]int{}
	for i := 0; i <= len(s)-L; i++ {
		x = (x<<2 | codeMap[s[i+L-1]]) & (1<<(L*2) - 1) // 从右边添加字符串
		cnt[x]++                                        // 记录这10位字符串出现次数
		if cnt[x] == 2 {
			res = append(res, s[i:i+L])
		}
		// x = x & ((1 << 20) - 1)
	}
	return res
}
