package WeeklyContest

func splitWordsBySeparator(words []string, separator byte) []string {
    var res []string
    for _, word := range words {
        begin := 0
        for i, w := range word {
            if w == int32(separator) {
                if i != begin {
                    res = append(res, word[begin:i])
                }
                begin = i + 1
            }
        }
        if begin < len(word) {
            res = append(res, word[begin:])
        }
    }
    return res
}
