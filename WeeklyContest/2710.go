package WeeklyContest

import "strings"

func removeTrailingZeros(num string) string {
    b := []byte(num)
    i := len(b)
    for ; i > 0 && num[i] == '0'; i-- {
    }
    return string(b[:i+1])
    
    return strings.TrimRight(num, "0") // 库函数写法
}
