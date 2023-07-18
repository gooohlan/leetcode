package WeeklyContest

func numberOfGoodSubarraySplits(nums []int) int {
    const mod = 1e9 + 7
    res, pre := 0, -1
    for i, num := range nums {
        if num > 0 {
            if pre >= 0 {
                res = res * (i - pre) % mod
            }
            pre = i
        }
    }
    if pre < 0 {
        return -1
    }
    return res
}
