package SlidingWindow

import "math"

var codeMap = map[byte]int{
	'A': 0,
	'C': 1,
	'G': 2,
	'T': 3,
}

func findRepeatedDnaSequences(s string) []string {
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
	for l, r := 0, 9; r < len(s); {
		x = x*R + codeMap[s[r]] // 从右边添加字符串
		r++
		cnt[x]++ // 记录这10位字符串出现次数
		if cnt[x] == 2 {
			res = append(res, s[l:r])
		}

		x = x - codeMap[s[l]]*RL
		l++
	}
	return res
}
