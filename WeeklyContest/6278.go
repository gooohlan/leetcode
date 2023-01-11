package WeeklyContest

func countDigits(num int) (res int) {
    for x := num; x > 0; x /= 10 {
        if num%(x%10) == 0 {
            res++
        }
    }
    return
}
