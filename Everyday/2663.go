package Everyday

func smallestBeautifulString(s string, k int) string {
    maxS := 'a' + byte(k) // 最大字符
    n := len(s)
    i := n - 1 // 从最后一个字符往前处理
    b := []byte(s)
    
    b[i]++
    for i < n {
        if b[i] == maxS { // 需要进位
            if i == 0 { // 遍历到第一个字符,没有结果
                return ""
            }
            
            b[i] = 'a'
            i--
            b[i]++ // 进位
        } else if i > 0 && b[i] == b[i-1] || i > 1 && b[i] == b[i-2] { // 左侧存在回文串
            b[i]++
        } else { // 左侧不存在回文串,往右寻找
            i++
        }
    }
    return string(b)
}
