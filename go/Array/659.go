package Array

import "fmt"

func isPossible(nums []int) bool {
    freq, need := make(map[int]int), make(map[int]int) // freq 剩余多少个未使用,need需要多少个
    
    for _, num := range nums {
        freq[num]++
    }
    
    for _, num := range nums {
        if freq[num] == 0 { // 已经用完
            continue
        }
        
        if need[num] > 0 {
            freq[num]-- // num可以放到当前子序列后面
            need[num]--
            need[num+1]++ // num+1所需+1
        } else if freq[num+1] > 0 && freq[num+2] > 0 {
            freq[num]--
            freq[num+1]--
            freq[num+2]--
            need[num+3]++
        } else {
            return false
        }
    }
    return true
}

func getPossible(nums []int) [][]int {
    freq, need := make(map[int]int), make(map[int][][]int) // freq 剩余多少个未使用,need需要多少个
    
    for _, num := range nums {
        freq[num]++
    }
    
    for _, num := range nums {
        if freq[num] == 0 { // 已经用完
            continue
        }
        
        if len(need[num]) > 0 {
            freq[num]--                        // num可以放到当前子序列后面
            seq := need[num][len(need[num])-1] // 取出最后一个
            seq = append(seq, num)
            need[num+1] = append(need[num+1], seq) // num+1所需+1
        } else if freq[num+1] > 0 && freq[num+2] > 0 {
            freq[num]--
            freq[num+1]--
            freq[num+2]--
            need[num+3] = append(need[num+3], []int{num, num + 1, num + 2})
        } else {
            return nil
        }
    }
    
    for _, v := range need {
        fmt.Println(v)
    }
    return nil
}
