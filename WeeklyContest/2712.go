package WeeklyContest

func minimumCost(s string) int64 {
    n := len(s)
    var sum int64
    for i := 1; i < n; i++ {
        if s[i] != s[i-1] {
            sum += int64(min(i, n-i))
        }
    }
    return sum
}
