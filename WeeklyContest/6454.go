package WeeklyContest

func makeSmallestPalindrome(s string) string {
    i, j := 0, len(s)-1
    str := []byte(s)
    for i < j {
        if str[i] != str[j] {
            if str[i] < str[j] {
                str[j] = str[i]
            } else {
                str[i] = str[j]
            }
        }
        i++
        j--
    }
    return string(str)
}
