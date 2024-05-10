package WeeklyContest

func minimizedStringLength(s string) int {
    m := map[int32]struct{}{}
    for _, i := range s {
        m[i] = struct{}{}
    }
    return len(m)
}
