package WeeklyContest


func xorBeauty(nums []int) int{
    res := 0
    for _, i := range nums {
        for _, j := range nums {
            for _, k := range nums {
                res ^= (i | j) & k
            }
        }
    }
    
    return res
}

// 上述答案不用看就知道会超时
// 我们假设nums = [a,b,c],根据上述暴力解法,一定会产生3^3=27个有效解
// (a | a) & a    (b | a) & a     (c | a) &  a
// (a | a) & b    (b | a) & b     (c | a) & b
// (a | a) & c    (b | a) & c     (c | a) & c
//
// (a | b) & a    (b | b) & a     (c | b) & a
// (a | b) & b    (b | b) & b     (c | b) & b
// (a | b) & c    (b | b) & c     (c | b) & c
//
// (a | c) & a    (b | c) & a     (c | c) & a
// (a | c) & b    (b | c) & b     (c | c) & a
// (a | c) & c    (b | c) & c     (c | c) & c

// 根据「按位或」的「对称性」，x | y = y | x,他们异或时为0,对答案没有贡献,所以上述只需要留下a=b的值

// (a | a) & a     (b | b) & a     (c | c) & a
// (a | a) & b     (b | b) & b     (c | c) & b
// (a | a) & c     (b | b) & c     (c | c) & c

// 又因为 a | a = a, a & a = a将上面式子简化一下
//   a       b & a     c & a
// a & b       b       c & b
// a & c     b & c       c

// 根据「按位与」的「对称性」，x & y = y & x,他们异或时为0,对答案没有贡献,最后剩下
// a    b   c
// 最后答案为a^b^c

func xorBeauty2(nums []int) int{
    res := 0
    for _, i := range nums {
        res ^= i
    }
    
    return res
}