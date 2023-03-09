package Backtrack

func generateParenthesis(n int) []string {
    if n == 0 {
        return nil
    }
    
    var (
        res       []string
        track     []byte
        backtrack func(left, right int)
    )
    
    backtrack = func(left, right int) {
        // 左括号比右括号多,不合法
        if right < left {
            return
        }
        // 数值小于0不合法
        if left < 0 || right < 0 {
            return
        }
        // 括号用完,加入结果集
        if left == 0 && right == 0 {
            res = append(res, string(track))
        }
        
        track = append(track, '(')
        backtrack(left-1, right)
        track = track[:len(track)-1]
        
        track = append(track, ')')
        backtrack(left, right-1)
        track = track[:len(track)-1]
    }
    backtrack(n, n)
    
    return res
}
