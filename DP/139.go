package DP

// 回溯解法
func wordBreak(s string, wordDict []string) bool {
    var found bool
    var track []string
    
    wordMap := make(map[string]bool)
    for _, word := range wordDict {
        wordMap[word] = true
    }
    
    var backtrack func(i int)
    backtrack = func(i int) {
        if found {
            return
        }
        
        if i == len(s) {
            found = true
            return
        }
        
        // for _, word := range wordDict {
        //     length := len(word)
        //     if i+length <= len(s) && s[i:i+length] == word {
        //         track = append(track, word)
        //         backtrack(i + length)
        //         track = track[:len(track)-1]
        //     }
        // }
        for length := 1; i+length <= len(s); length++ {
            prefix := s[i : i+length]
            if wordMap[prefix] {
                track = append(track, prefix)
                backtrack(i + length)
                track = track[:len(track)-1]
            }
        }
    }
    
    backtrack(0)
    return found
}
