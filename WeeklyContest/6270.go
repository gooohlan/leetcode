package WeeklyContest

func takeCharacters(s string, k int) int {
    n := len(s)
    // 先从右侧找到所有的字符
    arr := [3]int{}
    j := n
    for arr[0] < k || arr[1] < k || arr[2] < k {
        if j == 0 {
            return -1
        }
        j--
        arr[s[j]-'a']++
    }
    
    res := n - j // 左侧没有可用字符
    // 双指针,左边找到的时候,右侧处于同样字符时想右移
    for i := 0; i < n; i++ {
        arr[s[i]-'a']++
        for j < n && arr[s[j]-'a'] > k { // s[i]和s[j]相同时,j右移
            arr[s[j]-'a']--
            j++
        }
        res = min(res, i+1+n-j)
    }
    return res
}
