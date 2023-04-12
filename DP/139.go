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

func wordBreak2(s string, wordDict []string) bool {
    wordMap := make(map[string]bool)
    for _, word := range wordDict {
        wordMap[word] = true
    }
    // 备忘录，-1 代表未计算，0 代表无法凑出，1 代表可以凑出
    memo := make([]int, len(s))
    for i := 0; i < len(memo); i++ {
        memo[i] = -1
    }
    
    var dp func(i int) bool
    
    dp = func(i int) bool {
        if i == len(s) {
            return true
        }
        
        if memo[i] != -1 {
            if memo[i] == 0 {
                return false
            }
            return true
        }
        
        for j := 1; i+j <= len(s); j++ {
            prefix := s[i : i+j]
            if wordMap[prefix] {
                subProblem := dp(i + j)
                if subProblem {
                    memo[i] = 1
                    return true
                }
            }
        }
        memo[i] = 0
        return false
    }
    return dp(0)
}
