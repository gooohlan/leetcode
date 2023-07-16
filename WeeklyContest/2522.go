package WeeklyContest

func minimumPartition(s string, k int) int {
    res, num := 1, 0 // res初始化为1是因为最后结束循环时差一次自增
    for _, str := range s {
        v := int(str - '0')
        if v > k {
            return -1
        }
        num = num*10 + v
        if num > k {
            res++
            num = v
        }
    }
    return res
}
