package WeeklyContest

func longestValidSubstring(word string, forbidden []string) int {
    m := make(map[string]bool, len(forbidden))
    for _, s := range forbidden {
        m[s] = true
    }
    
    left := 0
    res := 0
    for right := range word {
        for i := right; i >= left && i > right-10; i-- { // forbidden中子串长度小于10
            if m[word[i:right+1]] {
                left = i + 1
                break
            }
        }
        res = max(res, right-left+1)
    }
    return res
}
