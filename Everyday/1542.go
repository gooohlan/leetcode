package Everyday

func longestAwesome(s string) int {
    const D = 10
    n := len(s)
    pos := [1 << D]int{}
    for i := range pos {
        pos[i] = n // 没有前缀和默认用最大值
    }

    pos[0] = -1
    pre := 0 // 因为是遍历的同时计算,所以pre即pre[j]
    res := 0
    for i, c := range s {
        pre ^= 1 << (c - '0')
        for d := 0; d < D; d++ {
            res = max(res, i-pos[pre^(1<<d)]) // 奇数pre[i]=pre[j]^(1<<d)
        }
        res = max(res, i-pos[pre]) // 偶数 pre[i]=pre[j]
        if pos[pre] == n {         // 标记前缀和首次出现位置
            pos[pre] = i
        }
    }
    return res
}
