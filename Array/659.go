package Array

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
