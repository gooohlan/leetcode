package Array

import (
	"math"
	"math/rand"
)

// 假设我们有数组 [0, n),blacklist中有m个数
// 我们将[0, n-m-1]放入非黑名单的数, [n-m, n-1]中存储黑名单的数, 然后随机返回[0, n-m-1]中的数即可
// 因为我们这里没有这个数组, 所以对于[0, n-m-1]范围内的黑名单数, 我们可以映射到[n-m, n-1]中非黑名单的数,
// 随机取出[0, n-m-1]中的x, x 分两种情况:
// x 非黑名单,直接返回
// x 是黑名单, 返回映射的非黑名单数
type Solution710 struct {
	mapping map[int]int
	sz      int
}

func Constructor710(n int, blacklist []int) Solution710 {
	mapping := make(map[int]int, 0)

	sz := n - len(blacklist)
	// 先将所有黑名单数字加入 map
	for _, i := range blacklist {
		if i >= sz { // i 本身就在黑名单数组范围内,可以不用添加
			mapping[i] = math.MaxInt
		}
	}

	last := n - 1
	for _, i := range blacklist {
		if i >= sz { // i 本身就在黑名单数组范围内,直接跳过
			continue
		}
		// 跳过所有黑名单中的数字
		for mapping[last] > 0 {
			last--
		}
		// 将黑名单中的索引映射到合法数字
		mapping[i] = last
		last--
	}

	return Solution710{mapping, sz}
}

func (s *Solution710) Pick() int {
	x := rand.Intn(s.sz)
	if s.mapping[x] > 0 {
		return s.mapping[x]
	}
	return x
}
