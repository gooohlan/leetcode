package Everyday

func removeTrailingZeros(num string) string {
    b := []byte(num)
    for i := len(b) - 1; i >= 0; i-- {
        if b[i] != '0' {
            return string(b[:i+1])
        }
    }
    return num
}
