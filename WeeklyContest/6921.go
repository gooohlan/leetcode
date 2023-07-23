package WeeklyContest

func splitWordsBySeparator(words []string, separator byte) []string {
    var res []string
    for _, word := range words {
        j := 0
        for i, w := range word {
            if w == int32(separator) {
                if i != j {
                    res = append(res, word[j:i])
                }
                j = i + 1
            }
        }
        if j < len(word) {
            res = append(res, word[j:])
        }
        // split := strings.Split(word, string(separator))
        // res = append(res, split...)
    }
    return res
}
