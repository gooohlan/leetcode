package WeeklyContest

import "fmt"

func minimizedStringLength(s string) int {
    var b []byte
    for i := 0; i < len(s); i++ {
        if len(b) > 0 {
            if b[len(b)-1] == s[i] {
                b = b[:len(b)-1]
                continue
            }
        }
        b = append(b, s[i])
    }
    fmt.Println(string(b))
    return len(b)
}
