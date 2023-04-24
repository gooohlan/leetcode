package DP

import "strings"

// 回溯1
func wordBreak140(s string, wordDict []string) []string {
    var res []string
    var track []string
    
    var backtrack func(i int)
    backtrack = func(i int) {
        // 当 i 超出字符串 s 的长度，说明找到了一种拼接方案
        if i == len(s) {
            res = append(res, strings.Join(track, " "))
            return
        }
        
        for _, word := range wordDict {
            length := len(word)
            if i+length <= len(s) && s[i:i+length] == word {
                track = append(track, word)
                backtrack(i + length)
                track = track[:len(track)-1]
            }
        }
        
    }
    backtrack(0)
    return res
}

// 回溯2
func wordBreak1402(s string, wordDict []string) []string {
    var res []string
    var track []string
    
    wordMap := make(map[string]bool)
    for _, word := range wordDict {
        wordMap[word] = true
    }
    
    var backtrack func(i int)
    backtrack = func(i int) {
        // 当 i 超出字符串 s 的长度，说明找到了一种拼接方案
        if i == len(s) {
            res = append(res, strings.Join(track, " "))
            return
        }
        
        for length := 1; length+i <= len(s); length++ {
            prefix := s[i : i+length]
            if wordMap[prefix] {
                track = append(track, prefix)
                backtrack(i + length)
                track = track[:len(track)-1]
            }
        }
    }
    backtrack(0)
    return res
}
