package WeeklyContest

func smallestString(s string) string {
    b := []byte(s)
    for i, c := range s {
        if c != 'a' {
            for ; i < len(b) && b[i] != 'a'; i++ {
                b[i]--
            }
            return string(b)
        }
    }
    b[len(b)-1] = 'z'
    return string(b)
}
