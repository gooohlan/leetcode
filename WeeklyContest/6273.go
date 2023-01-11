package WeeklyContest

// 先找到起始位置(值为0),再向左或向右寻找,找到-1的值,完成此次移动
func captureForts(forts []int) int {
    var res int
    for i, x := range forts {
        if x != 1 {
            continue
        }
        
        j := i - 1
        for j >= 0 && forts[j] == 0 {
            j--
        }
        if j >= 0 && forts[j] < 0 {
            res = max(res, i-j-1)
        }
        
        j = i + 1
        for j < len(forts) && forts[j] == 0 {
            j++
        }
        if j < len(forts) && forts[j] < 0 {
            res = max(res, j-i-1)
        }
    }
    return res
}
