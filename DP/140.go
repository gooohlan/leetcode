package DP

import (
    "container/list"
    "strings"
)

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

func wordBreak140DP(s string, wordDict []string) []string {
    wordMap := make(map[string]bool)
    for _, word := range wordDict {
        wordMap[word] = true
    }
    // 备忘录，-1 代表未计算，0 代表无法凑出，1 代表可以凑出
    memo := make([][]string, len(s))
    
    var dp func(i int) []string
    dp = func(i int) []string {
        if i == len(s) { // 区分匹配结束还是没有匹配结束标识
            return []string{""}
        }
        
        if len(memo[i]) != 0 {
            return memo[i]
        }
        
        var res []string
        for length := 1; length+i <= len(s); length++ {
            prefix := s[i : i+length]
            if wordMap[prefix] {
                subProblem := dp(i + length)
                
                for _, sub := range subProblem {
                    if sub == "" {
                        res = append(res, prefix)
                    } else {
                        res = append(res, prefix+" "+sub)
                    }
                    
                }
            }
        }
        memo[i] = res
        return res
    }
    return dp(0)
}

func wordBreak140DP2(s string, wordDict []string) []string {
    wordMap := make(map[string]bool)
    for _, word := range wordDict {
        wordMap[word] = true
    }
    // 备忘录，-1 代表未计算，0 代表无法凑出，1 代表可以凑出
    memo := make([]*list.List, len(s))
    
    var dp func(i int) *list.List
    dp = func(i int) *list.List {
        if i == len(s) {
            res := list.New()
            res.PushBack("")
            return res
        }
        
        if memo[i] != nil {
            return memo[i]
        }
        
        res := list.New()
        for length := 1; length+i <= len(s); length++ {
            sub := s[i : i+length]
            if wordMap[sub] {
                subProblem := dp(i + length)
                for e := subProblem.Front(); e != nil; e = e.Next() {
                    str := e.Value.(string)
                    if str == "" {
                        res.PushBack(sub)
                    } else {
                        res.PushBack(sub + " " + str)
                    }
                }
            }
        }
        memo[i] = res
        return res
    }
    var res []string
    for e := dp(0).Front(); e != nil; e = e.Next() {
        res = append(res, e.Value.(string))
    }
    return res
}
