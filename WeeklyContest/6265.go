package WeeklyContest

import (
    "sort"
    "strings"
)

func similarPairs(words []string) int {
    str := make([]string, len(words))
    for i, word := range words {
        split := strings.Split(word, "")
        sort.Strings(split)
        str[i] = split[0]
        for j := 1; j < len(split); j++ {
            if split[j] != split[j-1] {
                str[i] += split[j]
            }
        }
    }
    
    res := 0
    for i := 0; i < len(words); i++ {
        for j := i + 1; j < len(words); j++ {
            if str[i] == str[j] {
                res++
            }
        }
    }
    return res
}
