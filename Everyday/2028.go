package Everyday

func missingRolls(rolls []int, mean int, n int) []int {
    m := len(rolls)
    s := mean * (m + n)
    for _, roll := range rolls {
        s -= roll
    }
    
    if s < n || s > n*6 {
        return nil
    }
    
    avg := s / n
    poor := s % n
    ans := make([]int, n)
    for i := 0; i < n; i++ {
        ans[i] = avg
        if i < poor {
            ans[i] += 1
        }
    }
    
    return ans
}
