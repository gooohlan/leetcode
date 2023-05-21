package WeeklyContest

func minLength(s string) int {
    stack := make([]rune, 0)
    for _, char := range s {
        if len(stack) > 0 && ((stack[len(stack)-1] == 'A' && char == 'B') || (stack[len(stack)-1] == 'C' && char == 'D')) {
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, char)
        }
    }
    return len(stack)
}
