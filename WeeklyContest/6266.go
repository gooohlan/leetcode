package WeeklyContest

// 不断循环,计算n的质因数之和s,如果s==n说明无法再继续减小n了,返回n,否则更新n为s,继续循环

func smallestValue(n int) int {
    for {
        x, s := n, 0
        for i := 2; i*i <= x; i++ {
            for ; x%i == 0; x /= i {
                s += i
            }
        }
        if x > 1 {
            s += x
        }
        if s == n {
            return n
        }
        n = s
    }
}
