package WeeklyContest

func numberOfEmployeesWhoMetTarget(hours []int, target int) (res int) {
    for _, hour := range hours {
        if hour >= target {
            res++
        }
    }
    return
}
