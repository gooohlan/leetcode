package DataStructure

import "strconv"

func diffWaysToCompute(input string) []int {
    var res []int
    for i := 0; i < len(input); i++ {
        c := input[i]
        // 扫描算式 input 中的运算符
        if c == '-' || c == '*' || c == '+' {
            left := diffWaysToCompute(input[:i])
            right := diffWaysToCompute(input[i+1:])
            
            for _, l := range left {
                for _, r := range right {
                    if c == '+' {
                        res = append(res, l+r)
                    } else if c == '-' {
                        res = append(res, l-r)
                    } else if c == '*' {
                        res = append(res, l*r)
                    }
                }
            }
        }
    }
    
    if len(res) == 0 {
        num, _ := strconv.Atoi(input)
        res = append(res, num)
    }
    return res
}
