package WeeklyContest

func checkArray(nums []int, k int) bool {
    n := len(nums)
    d := make([]int, n+1)
    sumD := 0
    for i, num := range nums {
        sumD += d[i]
        num += sumD
        if num == 0 { // 已经为0,无需操作
            continue
        }
        if num < 0 || i+k > n { // 无法继续操作
            return false
        }
        sumD -= num // d[i] -= num
        d[i+k] += num
    }
    return true
}
