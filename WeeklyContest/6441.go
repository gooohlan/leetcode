package WeeklyContest

func punishmentNumber(n int) int {
    var sum int
    for i := 1; i <= n; i++ {
        if isPunishmentNumber(i) {
            sum += i * i
        }
    }
    return sum
}

func isPunishmentNumber(num int) bool {
    square := num * num
    // var res []int
    var backtrack func(int, int) int
    
    backtrack = func(r, sum int) int {
        i := 10
        if r == sum {
            return sum
        }
        for r/i > 0 {
            x := r % i
            if backtrack(r/i, sum-x) != -1 {
                return sum
            }
            i *= 10
        }
        return -1
    }
    
    sum := backtrack(square, num)
    return sum == num
}
